package app

import (
	"context"
	"time"

	"gchat/business/user/domain/repo"
	"gchat/pkg/protocol/pb"
)

type userApp struct{}

var UserApp = new(userApp)

func (*userApp) Get(ctx context.Context, userId int64) (*pb.User, error) {
	user, err := repo.UserRepo.Get(userId)
	if user == nil || err != nil {
		return nil, err
	}
	return user.ToProto(), err
}

func (*userApp) Update(ctx context.Context, userId int64, req *pb.UpdateUserReq) error {
	u, err := repo.UserRepo.Get(userId)
	if err != nil {
		return err
	}
	if u == nil {
		return nil
	}

	u.Nickname = req.Nickname
	u.Gender = req.Gender
	u.AvatarUrl = req.AvatarUrl
	u.Extra = req.Extra
	u.UpdateTime = time.Now()

	err = repo.UserRepo.Save(u)
	if err != nil {
		return err
	}
	return nil
}

func (*userApp) GetByIds(ctx context.Context, userIds []int64) (map[int64]*pb.User, error) {
	users, err := repo.UserRepo.GetByIds(userIds)
	if err != nil {
		return nil, err
	}

	pbUsers := make(map[int64]*pb.User, len(users))
	for i := range users {
		pbUsers[users[i].Id] = users[i].ToProto()
	}
	return pbUsers, nil
}
