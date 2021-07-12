package server

import (
	"time"

	v1 "github.com/dayu-go/comment/api/comment/service/v1"
	"github.com/dayu-go/comment/app/comment/service/internal/config"
	"github.com/dayu-go/comment/app/comment/service/internal/service"
	"github.com/dayu-go/gkit/log"
	"github.com/dayu-go/gkit/transport/grpc"
)

func NewGRPCServer(c config.Server, logger log.Logger, s *service.CommentService) *grpc.Server {
	var opts []grpc.ServerOption
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout > 0 {
		opts = append(opts, grpc.Timeout(time.Duration(c.Grpc.Timeout)*time.Second))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterCommentServiceServer(srv, s)
	return srv
}
