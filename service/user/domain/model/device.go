package model

type Device struct {
	Type   int32  // device type,1:Android；2：IOS；3：Windows; 4：MacOS；5：Web
	Token  string // device token
	Expire int64  // expiration time
}
