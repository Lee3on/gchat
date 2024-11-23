package gerrors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUnknown      = status.New(codes.Unknown, "Server error").Err() // Unknown server error
	ErrUnauthorized = newError(10000, "Please log in again")
	ErrBadRequest   = newError(10001, "Invalid request parameters")

	ErrBadCode         = newError(10010, "Invalid verification code")
	ErrNotInGroup      = newError(10011, "User is not in the group")
	ErrGroupNotExist   = newError(10013, "Group does not exist")
	ErrDeviceNotExist  = newError(10014, "Device does not exist")
	ErrAlreadyIsFriend = newError(10015, "The other party is already a friend")
	ErrUserNotFound    = newError(10016, "User not found")
)

func newError(code int, message string) error {
	return status.New(codes.Code(code), message).Err()
}
