package httpserver

import (
	v1 "github.com/dayu-go/comment/api/comment/service/v1"
	"github.com/dayu-go/comment/app/comment/service/internal/service"
	"github.com/dayu-go/comment/pkg/errs"
	"github.com/dayu-go/gkit/log"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	cs  *service.CommentService
	log *log.Helper
}

func NewCommentHandler(s *service.CommentService, l log.Logger) *CommentHandler {
	return &CommentHandler{
		cs:  s,
		log: log.NewHelper(l),
	}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	var req v1.CreateCommentRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Errorf("create comment parse request params failed, error:%s", err.Error())
		responseJSON(c, errs.ErrInvalidParam)
		return
	}
	_, err := h.cs.CreateComment(c, &req)
	if err != nil {
		log.Errorf("create comment failed, error:%s", err.Error())
		responseJSON(c, errs.ErrInvalidParam)
		return
	}
	responseJSON(c, nil, "hello")
}

func (h *CommentHandler) GetComment(c *gin.Context) {
	responseJSON(c, errs.ErrInvalidParam)
}

func (h *CommentHandler) GetReply(c *gin.Context) {
	responseJSON(c, nil, "hello")
}
