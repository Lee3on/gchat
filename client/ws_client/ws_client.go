package ws_client

import (
	"context"
	"fmt"
	"gchat/client/user"
	"gchat/pkg/protocol/pb"
	"gchat/pkg/util"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type WSClient struct {
	UserId   int64
	DeviceId int64
	Seq      int64
	Conn     *websocket.Conn
}

func (c *WSClient) Login() {
	resp, err := user.GetUserExtClient().SignIn(c.getCtx(), &pb.SignInReq{
		PhoneNumber: fmt.Sprintf("%d", c.UserId),
		Code:        "0",
		DeviceId:    c.DeviceId,
	})
	if err != nil {
		fmt.Println(err)
	}
	c.UserId = resp.UserId
}

func (c *WSClient) Start() {
	host := "127.0.0.1"
	if os.Getenv("WS_HOST") != "" {
		host = os.Getenv("WS_HOST")
	}
	conn, resp, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:8002/ws", host), http.Header{})
	if err != nil {
		fmt.Println("dial error", err)
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
	c.Conn = conn

	c.SignIn()
	c.SyncTrigger()
	go c.Heartbeat()
	go c.Receive()
}

func (c *WSClient) Output(pt pb.PackageType, requestId int64, message proto.Message) {
	var input = pb.Input{
		Type:      pt,
		RequestId: requestId,
	}
	if message != nil {
		bytes, err := proto.Marshal(message)
		if err != nil {
			fmt.Println(err)
			return
		}

		input.Data = bytes
	}

	writeBytes, err := proto.Marshal(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.Conn.WriteMessage(websocket.BinaryMessage, writeBytes)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *WSClient) SignIn() {
	signIn := pb.SignInInput{
		UserId:   c.UserId,
		DeviceId: c.DeviceId,
		Token:    "0",
	}
	c.Output(pb.PackageType_PT_SIGN_IN, time.Now().UnixNano(), &signIn)
}

func (c *WSClient) SyncTrigger() {
	c.Output(pb.PackageType_PT_SYNC, time.Now().UnixNano(), &pb.SyncInput{Seq: c.Seq})
}

func (c *WSClient) Heartbeat() {
	ticker := time.NewTicker(time.Minute * 4)
	for range ticker.C {
		c.Output(pb.PackageType_PT_HEARTBEAT, time.Now().UnixNano(), nil)
		fmt.Println("Heartbeat sent")
	}
}

func (c *WSClient) Receive() {
	for {
		_, bytes, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		c.HandlePackage(bytes)
	}
}

func (c *WSClient) HandlePackage(bytes []byte) {
	var output pb.Output
	err := proto.Unmarshal(bytes, &output)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch output.Type {
	case pb.PackageType_PT_HEARTBEAT:
		fmt.Println("Heartbeat response")
	case pb.PackageType_PT_SYNC:
		fmt.Println("Offline message sync started------")
		syncResp := pb.SyncOutput{}
		err := proto.Unmarshal(output.Data, &syncResp)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Offline message sync response code:", output.Code, "message:", output.Message)
		fmt.Printf("%+v \n", &output)
		for _, msg := range syncResp.Messages {
			log.Println(util.MessageToString(msg))
			c.Seq = msg.Seq
		}

		ack := pb.MessageACK{
			DeviceAck:   c.Seq,
			ReceiveTime: util.UnixMilliTime(time.Now()),
		}
		c.Output(pb.PackageType_PT_MESSAGE, output.RequestId, &ack)
		fmt.Println("Offline message sync finished------")
	case pb.PackageType_PT_MESSAGE:
		msg := pb.Message{}
		err := proto.Unmarshal(output.Data, &msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		log.Println("Message:", util.MessageToString(&msg))

		c.Seq = msg.Seq
		ack := pb.MessageACK{
			DeviceAck:   msg.Seq,
			ReceiveTime: util.UnixMilliTime(time.Now()),
		}
		c.Output(pb.PackageType_PT_MESSAGE, output.RequestId, &ack)
	default:
		fmt.Println("switch other")
	}
}

func (c *WSClient) getCtx() context.Context {
	token := "0"
	return metadata.NewOutgoingContext(context.TODO(), metadata.Pairs(
		"user_id", fmt.Sprintf("%d", c.UserId),
		"device_id", fmt.Sprintf("%d", c.DeviceId),
		"token", token,
		"request_id", strconv.FormatInt(time.Now().UnixNano(), 10)))
}
