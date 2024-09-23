package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/Yuki-TU/dynamodb-sample/config"
	"github.com/Yuki-TU/dynamodb-sample/controller"
	"github.com/Yuki-TU/dynamodb-sample/gen"
	"github.com/Yuki-TU/dynamodb-sample/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(context.Background()); err != nil {
		slog.Error(fmt.Sprintf("failed to terminated server: %s", err.Error()))
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	if err := config.Load(); err != nil {
		return err
	}

	r := gin.Default()
	repo, err := repository.New(ctx)
	if err != nil {
		return err
	}

	// ルーティング初期化
	impl := controller.NewController(repo)
	strictServer := gen.NewStrictHandler(impl, nil)
	gen.RegisterHandlersWithOptions(
		r,
		strictServer,
		gen.GinServerOptions{BaseURL: "/api/"},
	)

	// サーバー起動
	log.Printf("Listening and serving HTTP on :%v", config.Get().Port)
	server := NewServer(r, fmt.Sprintf(":%d", config.Get().Port))
	return server.Run(ctx)
}
