package connect

import (
	"container/list"
	"context"
	"gchat/config"
	"gchat/pkg/grpclib"
	"gchat/pkg/logger"
	"gchat/pkg/protocol/pb"
	"gchat/pkg/rpc"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const (
	CoonTypeTCP int8 = 1 // TCP connection
	ConnTypeWS  int8 = 2 // WebSocket connection
)

type Conn struct {
	CoonType int8            // Connection type
	TCP      net.Conn        // TCP connection
	WSMutex  sync.Mutex      // WebSocket write lock
	WS       *websocket.Conn // WebSocket connection
	UserId   int64           // User ID
	DeviceId int64           // Device ID
	RoomId   int64           // Subscribed room ID
	Element  *list.Element   // Linked list node
}

// Write sends data to the connection.
func (c *Conn) Write(bytes []byte) error {
	if c.CoonType == CoonTypeTCP {
		return c.WriteToTCP(bytes)
	} else if c.CoonType == ConnTypeWS {
		return c.WriteToWS(bytes)
	}
	logger.Logger.Error("unknown conn type", zap.Any("conn", c))
	return nil
}

// WriteToTCP writes data to a TCP connection.
func (c *Conn) WriteToTCP(bytes []byte) error {
	_, err := c.TCP.Write(bytes)
	return err
}

// WriteToWS writes data to a WebSocket connection.
func (c *Conn) WriteToWS(bytes []byte) error {
	c.WSMutex.Lock()
	defer c.WSMutex.Unlock()

	err := c.WS.SetWriteDeadline(time.Now().Add(10 * time.Millisecond))
	if err != nil {
		return err
	}
	return c.WS.WriteMessage(websocket.BinaryMessage, bytes)
}

// Close closes the connection.
func (c *Conn) Close() error {
	// Remove the mapping between device and connection.
	if c.DeviceId != 0 {
		DeleteConn(c.DeviceId)
	}

	// Unsubscribe asynchronously to avoid deadlock.
	go func() {
		SubscribedRoom(c, 0)
	}()

	if c.DeviceId != 0 {
		_, _ = rpc.GetLogicIntClient().Offline(context.TODO(), &pb.OfflineReq{
			UserId:     c.UserId,
			DeviceId:   c.DeviceId,
			ClientAddr: c.GetAddr(),
		})
	}

	if c.CoonType == CoonTypeTCP {
		return c.TCP.Close()
	} else if c.CoonType == ConnTypeWS {
		return c.WS.Close()
	}
	return nil
}

// GetAddr returns the connection address.
func (c *Conn) GetAddr() string {
	if c.CoonType == CoonTypeTCP {
		return c.TCP.RemoteAddr().String()
	} else if c.CoonType == ConnTypeWS {
		return c.WS.RemoteAddr().String()
	}
	return ""
}

// HandleMessage processes incoming messages.
func (c *Conn) HandleMessage(bytes []byte) {
	var input = new(pb.Input)
	err := proto.Unmarshal(bytes, input)
	if err != nil {
		logger.Logger.Error("unmarshal error", zap.Error(err), zap.Int("len", len(bytes)))
		return
	}
	logger.Logger.Debug("HandleMessage", zap.Any("input", input))

	// Intercept unauthenticated users
	if input.Type != pb.PackageType_PT_SIGN_IN && c.UserId == 0 {
		// Notify the user they are not logged in
		logger.Logger.Info("user not logged in")
		return
	}

	switch input.Type {
	case pb.PackageType_PT_SIGN_IN:
		c.SignIn(input)
	case pb.PackageType_PT_SYNC:
		c.Sync(input)
	case pb.PackageType_PT_HEARTBEAT:
		c.Heartbeat(input)
	case pb.PackageType_PT_MESSAGE:
		c.MessageACK(input)
	case pb.PackageType_PT_SUBSCRIBE_ROOM:
		c.SubscribedRoom(input)
	default:
		logger.Logger.Error("handler switch other")
	}
}

// Send sends a message to the client.
func (c *Conn) Send(pt pb.PackageType, requestId int64, message proto.Message, err error) {
	var output = pb.Output{
		Type:      pt,
		RequestId: requestId,
	}

	if err != nil {
		status, _ := status.FromError(err)
		output.Code = int32(status.Code())
		output.Message = status.Message()
	}

	if message != nil {
		msgBytes, err := proto.Marshal(message)
		if err != nil {
			logger.Sugar.Error(err)
			return
		}
		output.Data = msgBytes
	}

	outputBytes, err := proto.Marshal(&output)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	err = c.Write(outputBytes)
	if err != nil {
		logger.Sugar.Error(err)
		c.Close()
		return
	}
}

// SignIn processes a login request.
func (c *Conn) SignIn(input *pb.Input) {
	var signIn pb.SignInInput
	err := proto.Unmarshal(input.Data, &signIn)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Logger.Debug("SignIn", zap.Any("signIn", signIn))

	_, err = rpc.GetLogicIntClient().ConnSignIn(grpclib.ContextWithRequestId(context.TODO(), input.RequestId), &pb.ConnSignInReq{
		UserId:     signIn.UserId,
		DeviceId:   signIn.DeviceId,
		Token:      signIn.Token,
		ConnAddr:   config.Config.ConnectLocalAddr,
		ClientAddr: c.GetAddr(),
	})

	c.Send(pb.PackageType_PT_SIGN_IN, input.RequestId, nil, err)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	c.UserId = signIn.UserId
	c.DeviceId = signIn.DeviceId
	SetConn(signIn.DeviceId, c)
}

// Sync synchronizes messages.
func (c *Conn) Sync(input *pb.Input) {
	var sync pb.SyncInput
	err := proto.Unmarshal(input.Data, &sync)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	resp, err := rpc.GetLogicIntClient().Sync(grpclib.ContextWithRequestId(context.TODO(), input.RequestId), &pb.SyncReq{
		UserId:   c.UserId,
		DeviceId: c.DeviceId,
		Seq:      sync.Seq,
	})

	var message proto.Message
	if err == nil {
		message = &pb.SyncOutput{Messages: resp.Messages, HasMore: resp.HasMore}
	}
	c.Send(pb.PackageType_PT_SYNC, input.RequestId, message, err)
}

// Heartbeat processes a heartbeat message.
func (c *Conn) Heartbeat(input *pb.Input) {
	c.Send(pb.PackageType_PT_HEARTBEAT, input.RequestId, nil, nil)
	logger.Sugar.Infow("heartbeat", "device_id", c.DeviceId, "user_id", c.UserId)
}

// MessageACK processes a message acknowledgment.
func (c *Conn) MessageACK(input *pb.Input) {
	var messageACK pb.MessageACK
	err := proto.Unmarshal(input.Data, &messageACK)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	_, _ = rpc.GetLogicIntClient().MessageACK(grpclib.ContextWithRequestId(context.TODO(), input.RequestId), &pb.MessageACKReq{
		UserId:      c.UserId,
		DeviceId:    c.DeviceId,
		DeviceAck:   messageACK.DeviceAck,
		ReceiveTime: messageACK.ReceiveTime,
	})
}

// SubscribedRoom processes room subscription.
func (c *Conn) SubscribedRoom(input *pb.Input) {
	var subscribeRoom pb.SubscribeRoomInput
	err := proto.Unmarshal(input.Data, &subscribeRoom)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	SubscribedRoom(c, subscribeRoom.RoomId)
	c.Send(pb.PackageType_PT_SUBSCRIBE_ROOM, input.RequestId, nil, nil)
	_, err = rpc.GetLogicIntClient().SubscribeRoom(context.TODO(), &pb.SubscribeRoomReq{
		UserId:   c.UserId,
		DeviceId: c.DeviceId,
		RoomId:   subscribeRoom.RoomId,
		Seq:      subscribeRoom.Seq,
		ConnAddr: config.Config.ConnectLocalAddr,
	})
	if err != nil {
		logger.Logger.Error("SubscribedRoom error", zap.Error(err))
	}
}
