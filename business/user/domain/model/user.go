package model

import (
	"gchat/pkg/protocol/pb"
	"time"
)

type User struct {
	Id          int64     // user id
	PhoneNumber string    // user phone number
	Nickname    string    // nickname
	Gender      int32     // gender, 0: prefer not to say; 1:male; 2:female; 3:non-binary; 4: other
	AvatarUrl   string    // user avatar url
	Extra       string    // extra info
	CreateTime  time.Time // create time
	UpdateTime  time.Time // update time
}

func (u *User) ToProto() *pb.User {
	if u == nil {
		return nil
	}

	return &pb.User{
		UserId:     u.Id,
		Nickname:   u.Nickname,
		Gender:     u.Gender,
		AvatarUrl:  u.AvatarUrl,
		Extra:      u.Extra,
		CreateTime: u.CreateTime.Unix(),
		UpdateTime: u.UpdateTime.Unix(),
	}
}
