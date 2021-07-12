package httpserver

import (
	"net/http"

	"github.com/dayu-go/gkit/errors"
	"github.com/gin-gonic/gin"
)

func responseJSON(ctx *gin.Context, err *errors.Error, data ...interface{}) {
	m := map[string]interface{}{
		"code": 0,
		"msg":  "success",
	}
	if len(data) > 0 && data[0] != nil {
		m["data"] = data[0]
	}
	if err != nil {
		m["code"] = errors.Code(err)
		m["msg"] = errors.Message(err)
	}
	ctx.JSON(http.StatusOK, m)
}
