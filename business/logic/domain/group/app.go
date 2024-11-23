package group

import (
	"context"
	"gchat/business/logic/domain/group/entity"
	"gchat/business/logic/domain/group/repo"
	"gchat/pkg/protocol/pb"
)

type app struct{}

var App = new(app)

func (*app) CreateGroup(ctx context.Context, userId int64, in *pb.CreateGroupReq) (int64, error) {
	group := entity.CreateGroup(userId, in)
	err := repo.GroupRepo.Save(group)
	if err != nil {
		return 0, err
	}
	return group.Id, nil
}

func (*app) GetGroup(ctx context.Context, groupId int64) (*pb.Group, error) {
	group, err := repo.GroupRepo.Get(groupId)
	if err != nil {
		return nil, err
	}

	return group.ToProto(), nil
}

// GetUserGroups gets the list of groups that the user has joined
func (*app) GetUserGroups(ctx context.Context, userId int64) ([]*pb.Group, error) {
	groups, err := repo.GroupUserRepo.ListByUserId(userId)
	if err != nil {
		return nil, err
	}

	pbGroups := make([]*pb.Group, len(groups))
	for i := range groups {
		pbGroups[i] = groups[i].ToProto()
	}
	return pbGroups, nil
}

func (*app) Update(ctx context.Context, userId int64, update *pb.UpdateGroupReq) error {
	group, err := repo.GroupRepo.Get(update.GroupId)
	if err != nil {
		return err
	}

	err = group.Update(ctx, update)
	if err != nil {
		return err
	}

	err = repo.GroupRepo.Save(group)
	if err != nil {
		return err
	}

	err = group.PushUpdate(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}

func (*app) AddMembers(ctx context.Context, userId, groupId int64, userIds []int64) ([]int64, error) {
	group, err := repo.GroupRepo.Get(groupId)
	if err != nil {
		return nil, err
	}
	existIds, addedIds, err := group.AddMembers(ctx, userIds)
	if err != nil {
		return nil, err
	}
	err = repo.GroupRepo.Save(group)
	if err != nil {
		return nil, err
	}

	err = group.PushAddMember(ctx, userId, addedIds)
	if err != nil {
		return nil, err
	}
	return existIds, nil
}

func (*app) UpdateMember(ctx context.Context, in *pb.UpdateGroupMemberReq) error {
	group, err := repo.GroupRepo.Get(in.GroupId)
	if err != nil {
		return err
	}
	err = group.UpdateMember(ctx, in)
	if err != nil {
		return err
	}
	err = repo.GroupRepo.Save(group)
	if err != nil {
		return err
	}
	return nil
}

func (*app) DeleteMember(ctx context.Context, groupId int64, userId int64, optId int64) error {
	group, err := repo.GroupRepo.Get(groupId)
	if err != nil {
		return err
	}
	err = group.DeleteMember(ctx, userId)
	if err != nil {
		return err
	}
	err = repo.GroupRepo.Save(group)
	if err != nil {
		return err
	}

	err = group.PushDeleteMember(ctx, optId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (*app) GetMembers(ctx context.Context, groupId int64) ([]*pb.GroupMember, error) {
	group, err := repo.GroupRepo.Get(groupId)
	if err != nil {
		return nil, err
	}
	return group.GetMembers(ctx)
}

func (*app) SendMessage(ctx context.Context, fromDeviceID, fromUserID int64, req *pb.SendMessageReq) (int64, error) {
	group, err := repo.GroupRepo.Get(req.ReceiverId)
	if err != nil {
		return 0, err
	}

	return group.SendMessage(ctx, fromDeviceID, fromUserID, req)
}
