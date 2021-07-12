package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/dayu-go/comment/app/comment/service/internal/store"
	"github.com/dayu-go/gkit/log"
)

type CommentBiz struct {
	store *store.Store
	log   *log.Helper
}

func NewCommentBiz(store *store.Store, logger log.Logger) *CommentBiz {
	return &CommentBiz{
		store: store,
		log:   log.NewHelper(logger.Fields(map[string]interface{}{"module": "biz/comment"})),
	}
}

type CommentRequest struct {
	UserId   int64
	ObjId    int64
	objType  int
	Root     int64
	IP       string
	Platform int
	Device   string
	Meta     string
	Message  string
}

func (cc *CommentBiz) CreateComment(ctx context.Context, c *CommentRequest) (resp *store.Comment, err error) {
	now := time.Now().Format(DateTimeFormat)

	// find obj

	// find root comment
	ok, err := cc.store.ExistsCommentIndex(ctx, c.Root)
	if err != nil {
		return
	}
	if !ok {
		err = fmt.Errorf("root comment index is not exists, id:%d", c.Root)
		return
	}

	// create comment
	newId, err := cc.store.CreateComment(ctx, store.CreateCommentRequest{
		ObjId:      c.ObjId,
		ObjType:    c.objType,
		UserId:     c.UserId,
		Root:       c.Root,
		IP:         c.IP,
		Platform:   c.Platform,
		Device:     c.Device,
		Meta:       c.Meta,
		Message:    c.Message,
		CreateTime: now,
		UpdateTime: now,
	})
	if err != nil {
		return
	}
	resp.Id = newId
	return
}

func (cc *CommentBiz) GetComment(ctx context.Context, id int64) (*store.Comment, error) {
	return nil, nil
}

func (cc *CommentBiz) UpdateComment(ctx context.Context, c *store.Comment) (*store.Comment, error) {
	return nil, nil
}

func (cc *CommentBiz) ListComment(ctx context.Context, pageNum, pageSize int64) ([]*store.Comment, error) {
	return nil, nil
}
