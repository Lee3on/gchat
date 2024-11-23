package grpclib

import (
	"context"
	"gchat/pkg/gerrors"
	"gchat/pkg/logger"
	"strconv"

	"google.golang.org/grpc/metadata"
)

const (
	CtxUserId    = "user_id"
	CtxDeviceId  = "device_id"
	CtxToken     = "token"
	CtxRequestId = "request_id"
)

func ContextWithRequestId(ctx context.Context, requestId int64) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(CtxRequestId, strconv.FormatInt(requestId, 10)))
}

func Get(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values, ok := md[key]
	if !ok || len(values) == 0 {
		return ""
	}
	return values[0]
}

// GetCtxRequestId get app_id from ctx
func GetCtxRequestId(ctx context.Context) int64 {
	requestIdStr := Get(ctx, CtxRequestId)
	requestId, err := strconv.ParseInt(requestIdStr, 10, 64)
	if err != nil {
		return 0
	}
	return requestId
}

// GetCtxData get user info from ctx
func GetCtxData(ctx context.Context) (int64, int64, error) {
	var (
		userId   int64
		deviceId int64
		err      error
	)

	userIdStr := Get(ctx, CtxUserId)
	userId, err = strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, 0, gerrors.ErrUnauthorized
	}

	deviceIdStr := Get(ctx, CtxDeviceId)
	deviceId, err = strconv.ParseInt(deviceIdStr, 10, 64)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, 0, gerrors.ErrUnauthorized
	}
	return userId, deviceId, nil
}

// GetCtxToken get token from ctx
func GetCtxToken(ctx context.Context) string {
	return Get(ctx, CtxToken)
}

// NewAndCopyRequestId create a ctx and copy request id
func NewAndCopyRequestId(ctx context.Context) context.Context {
	newCtx := context.TODO()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return newCtx
	}

	requestIds, ok := md[CtxRequestId]
	if !ok && len(requestIds) == 0 {
		return newCtx
	}
	return metadata.NewOutgoingContext(newCtx, metadata.Pairs(CtxRequestId, requestIds[0]))
}
