package httpserver

import (
	"github.com/dayu-go/comment/pkg/err"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	responseJSON(c, nil, "hello")
}

func GetComment(c *gin.Context) {
	responseJSON(c, err.ErrInvalidParam)
}

func GetReply(c *gin.Context) {
	responseJSON(c, nil, "hello")
}
