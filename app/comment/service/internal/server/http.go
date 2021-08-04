package server

import (
	"github.com/dayu-go/comment/app/comment/service/internal/config"
	"github.com/dayu-go/comment/app/comment/service/internal/httpserver"
	"github.com/dayu-go/gkit/log"
	"github.com/dayu-go/gkit/middleware/logging"
	"github.com/dayu-go/gkit/middleware/validate"
	"github.com/dayu-go/gkit/transport/http"
)

func NewHTTPServer(c config.Server, logger log.Logger, s *httpserver.CommentHandler) *http.Server {

	opts := []http.ServerOption{
		http.Middleware(
			logging.Server(logger),
			validate.Validator()),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout > 0 {
		opts = append(opts, http.Timeout(c.Http.Timeout))
	}

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", httpserver.RegisterHTTPServer(s))
	return srv
}
