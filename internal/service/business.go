package service

import (
	"context"
	"fmt"

	pb "review-b/api/business/v1"
	"review-b/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer

	uc  *biz.ReviewReplyUsecase
	log *log.Helper
}

func NewBusinessService(uc *biz.ReviewReplyUsecase, logger log.Logger) *BusinessService {
	return &BusinessService{uc: uc, log: log.NewHelper(logger)}
}

func (s *BusinessService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	s.log.Infof("ReplyReview: %+v", req)
	fmt.Printf("ReplyReview: %+v", req)
	resp, err := s.uc.CreateReviewReply(ctx, req)
	if err != nil {
		s.log.Errorf("ReplyReview err: %+v", err)
		return nil, err
	}
	s.log.Infof("ReplyReview success: %+v", resp)
	return resp, nil
}

func (s *BusinessService) AppealReview(ctx context.Context, req *pb.AppealReviewRequest) (*pb.AppealReviewReply, error) {
	s.log.Infof("AppealReview: %+v", req)
	resp, err := s.uc.CreateReviewAppeal(ctx, req)
	if err != nil {
		s.log.Errorf("AppealReview err: %+v", err)
		return nil, err
	}
	s.log.Infof("AppealReview success: %+v", resp)
	return resp, nil
}
