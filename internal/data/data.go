package data

import (
	"context"
	"review-b/internal/conf"

	v1 "review-b/api/review/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewGrpcClient, NewData, NewReviewReplyRepo)

// Data .
type Data struct {
	// grpc client
	grpcClient v1.ReviewClient
	logger     *log.Helper
}

// NewData .
func NewData(client v1.ReviewClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		grpcClient: client,
		logger:     log.NewHelper(logger),
	}, cleanup, nil
}

func NewGrpcClient(c *conf.Data, logger log.Logger) v1.ReviewClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9092"),
	)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	return v1.NewReviewClient(conn)
}
