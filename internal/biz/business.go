package biz

import (
	"context"
	pb "review-b/api/business/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type ReviewReplyRepo interface {
	CreateReviewReply(ctx context.Context, req *pb.ReplyReviewRequest) (resp *pb.ReplyReviewReply, err error)
	CreateReviewAppeal(ctx context.Context, req *pb.AppealReviewRequest) (resp *pb.AppealReviewReply, err error)
}

type ReviewReplyUsecase struct {
	repo ReviewReplyRepo
	log  *log.Helper
}

func NewReviewReplyUsecase(repo ReviewReplyRepo, logger log.Logger) *ReviewReplyUsecase {
	return &ReviewReplyUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *ReviewReplyUsecase) CreateReviewReply(ctx context.Context, req *pb.ReplyReviewRequest) (resp *pb.ReplyReviewReply, err error) {
	uc.log.Infof("SaveReviewReplyAndUpdateReview: %+v", req)
	resp, err = uc.repo.CreateReviewReply(ctx, req)
	if err != nil {
		uc.log.Errorf("CreateReviewReply err: %+v", err)
		return nil, err
	}
	uc.log.Infof("CreateReviewReply success: %+v", resp)
	return resp, nil
}

func (uc *ReviewReplyUsecase) CreateReviewAppeal(ctx context.Context, req *pb.AppealReviewRequest) (resp *pb.AppealReviewReply, err error) {
	uc.log.Infof("CreateReviewAppeal: %+v", req)
	resp, err = uc.repo.CreateReviewAppeal(ctx, req)
	if err != nil {
		uc.log.Errorf("CreateReviewAppeal err: %+v", err)
		return nil, err
	}
	uc.log.Infof("CreateReviewAppeal success: %+v", resp)
	return resp, nil
}

