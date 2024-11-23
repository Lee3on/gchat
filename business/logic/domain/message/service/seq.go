package service

import (
	"context"
	"gchat/business/logic/domain/message/repo"
)

type seqService struct{}

var SeqService = new(seqService)

// GetUserNext Get the next sequence number of the user
func (*seqService) GetUserNext(ctx context.Context, userId int64) (int64, error) {
	return repo.SeqRepo.Incr(repo.SeqObjectTypeUser, userId)
}
