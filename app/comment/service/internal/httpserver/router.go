package httpserver

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPServer(h *CommentHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/_ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// router.Use(http.CORSMiddleware())

	rg1 := router.Group("/api/v1")
	rg1.POST("/comment/add", h.CreateComment)
	rg1.GET("/comment/main", h.GetComment)
	rg1.GET("/comment/reply", h.GetReply)

	return router
}
