package service

import (
	"context"

	v1 "github.com/dayu-go/comment/api/comment/service/v1"
	"github.com/dayu-go/comment/app/comment/service/internal/biz"
	"github.com/dayu-go/gkit/log"
)

type CommentService struct {
	v1.UnimplementedCommentServiceServer
	cb  *biz.CommentBiz
	log *log.Helper
}

type CreateCommentRequest struct {
	UserId  int64
	ObjId   int64
	ObjType int
	RootId  int
	Content string
}

func (s *CommentService) CreateComment(ctx context.Context, req *v1.CreateCommentRequest) (res *v1.CreateCommentResponse, err error) {
	c, err := s.cb.CreateComment(ctx, &biz.CommentRequest{})
	return &v1.CreateCommentResponse{
		Id: c.Id,
	}, err
}

type Reply struct {
	RpID       int64  `json:"id"`
	ObjId      uint64 `json:"obj_id"`
	ObjType    int8   `json:"obj_type"`
	UserId     int64  `json:"userId"`
	Root       int64  `json:"root"`
	Parent     int64  `json:"parent"`
	Count      int    `json:"count"`
	RootCount  int    `json:"root_count"`
	Floor      int    `json:"floor"`
	State      int8   `json:"state"`
	Attr       int8   `json:"attr"`
	CreateTime int64  `json:"create_time"`
	Like       int    `json:"like"`
	// user info
	User *User `json:"user,omitempty"`
	// other
	Content *Content `json:"content,omitempty"`
	Replies []*Reply `json:"replies,omitempty"`
}

type User struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Sex    string `json:"sex"`
	Avatar string `json:"avatar"`
}

type Content struct {
	Message string
}

func (s *CommentService) GetComment(ctx context.Context, req *v1.GetCommentRequest) (*v1.GetCommentResponse, error) {
	return nil, nil
}

func (s *CommentService) GetReply(ctx context.Context, req *v1.GetReplyRequest) (*v1.GetReplyResponse, error) {
	return nil, nil
}
