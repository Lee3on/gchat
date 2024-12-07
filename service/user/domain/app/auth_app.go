package app

import (
	"context"
	"gchat/service/user/domain/service"
)

type authApp struct{}

var AuthApp = new(authApp)

// SignIn signs in the user (and keeps the long connection)
func (*authApp) SignIn(ctx context.Context, phoneNumber, code string, deviceId int64) (bool, int64, string, error) {
	return service.AuthService.SignIn(ctx, phoneNumber, code, deviceId)
}

// Auth checks if the user is authenticated
func (*authApp) Auth(ctx context.Context, userId, deviceId int64, token string) error {
	return service.AuthService.Auth(ctx, userId, deviceId, token)
}
