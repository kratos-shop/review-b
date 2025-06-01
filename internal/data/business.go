package data

import (
	"context"

	"review-b/internal/biz"

	pb "review-b/api/business/v1"
	v1 "review-b/api/review/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type reviewReplyRepo struct {
	data *Data
	log  *log.Helper
}

func NewReviewReplyRepo(data *Data, logger log.Logger) biz.ReviewReplyRepo {
	return &reviewReplyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewReplyRepo) CreateReviewReply(ctx context.Context, req *pb.ReplyReviewRequest) (resp *pb.ReplyReviewReply, err error) {
	r.log.Infof("CreateReviewReply: %+v", req)
	grpcResp, err := r.data.grpcClient.ReplyReview(ctx, &v1.ReplyReviewRequest{
		ReviewId:  req.ReviewId,
		StoreId:   req.StoreId,
		Content:   req.Content,
		PicInfo:   req.PicInfo,
		VideoInfo: req.VideoInfo,
	})
	if err != nil {
		r.log.Errorf("CreateReviewReply err: %+v", err)
		return nil, err
	}
	r.log.Infof("CreateReviewReply success: %+v", grpcResp)
	return &pb.ReplyReviewReply{Id: grpcResp.Id}, nil
}
