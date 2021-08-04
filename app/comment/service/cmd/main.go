package main

import (
	"github.com/dayu-go/comment/app/comment/service/internal/biz"
	"github.com/dayu-go/comment/app/comment/service/internal/config"
	"github.com/dayu-go/comment/app/comment/service/internal/httpserver"
	"github.com/dayu-go/comment/app/comment/service/internal/server"
	"github.com/dayu-go/comment/app/comment/service/internal/service"
	"github.com/dayu-go/comment/app/comment/service/internal/store"
	"github.com/dayu-go/gkit/app"
	"github.com/dayu-go/gkit/log"
)

func main() {

	// load config
	if err := config.Load(); err != nil {
		panic(err)
	}

	// log
	opts := []log.Option{}
	if config.Conf.App.Env == "prod" {
		opts = append(opts, log.WithLevel(log.LevelInfo))
	} else {
		opts = append(opts, log.WithLevel(log.LevelDebug))
	}
	logger := log.NewHelper(log.NewLogger(opts...))

	app, err := newApp(&config.Conf, logger)
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApp(c *config.Config, logger log.Logger) (*app.App, error) {
	store := store.NewStore().NewDB(c.DB.Dayu)
	commentBiz := biz.NewCommentBiz(store, logger)
	cs := service.NewCommentService(commentBiz, logger)
	handler := httpserver.NewCommentHandler(cs, logger)

	hs := server.NewHTTPServer(c.Server, logger, handler)
	// gs := server.NewGRPCServer(c.Server, logger, cs)

	return app.New(
		app.Name("comment"),
		app.Version("v1.0.0"),
		app.Server(hs),
	), nil
}
