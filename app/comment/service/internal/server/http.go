package server

import (
	"github.com/dayu-go/comment/app/comment/service/internal/config"
	"github.com/dayu-go/comment/app/comment/service/internal/httpserver"
	"github.com/dayu-go/gkit/log"
	"github.com/dayu-go/gkit/transport/http"
	"github.com/gin-gonic/gin"
)

func NewHTTPServer(c config.Server, logger log.Logger) *http.Server {

	// TODO: middleware
	opts := []http.ServerOption{}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout > 0 {
		opts = append(opts, http.Timeout(c.Http.Timeout))
	}

	router := gin.Default()
	router.GET("/_ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	rg1 := router.Group("/api/v1")
	rg1.POST("/comment/add", httpserver.CreateComment)
	rg1.GET("/comment/main", httpserver.GetComment)
	rg1.GET("/comment/reply", httpserver.GetReply)

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", router)
	return srv
}
