package connect

import (
	"context"
	"gchat/pkg/grpclib"
	"gchat/pkg/logger"
	"gchat/pkg/protocol/pb"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ConnIntServer struct {
	pb.UnsafeConnectIntServer
}

// DeliverMessage deliver message to device
func (s *ConnIntServer) DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*emptypb.Empty, error) {
	resp := &emptypb.Empty{}

	// Get the TCP connection corresponding to the device
	conn := GetConn(req.DeviceId)
	if conn == nil {
		logger.Logger.Warn("GetConn warn", zap.Int64("device_id", req.DeviceId))
		return resp, nil
	}

	// Check if the device_id in the request is the same as the device_id in the connection
	if conn.DeviceId != req.DeviceId {
		logger.Logger.Warn("GetConn warn", zap.Int64("device_id", req.DeviceId))
		return resp, nil
	}

	// Send message to device
	conn.Send(pb.PackageType_PT_MESSAGE, grpclib.GetCtxRequestId(ctx), req.Message, nil)
	return resp, nil
}
