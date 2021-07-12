package service

import (
	"github.com/dayu-go/comment/app/comment/service/internal/biz"
	"github.com/dayu-go/gkit/log"
)

func NewCommentService(cb *biz.CommentBiz, logger log.Logger) *CommentService {
	return &CommentService{
		cb:  cb,
		log: log.NewHelper(logger.Fields(map[string]interface{}{"module": "service/comment"})),
	}
}
