package cmd

import (
	"context"
	"log"
	"net/http"
	"todo-list-app/database"
	"todo-list-app/domain"
	middleware "todo-list-app/internal/http/midlleware"
	"todo-list-app/internal/repository"
	"todo-list-app/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	UserUsecase domain.UserUsecase
}

func Init() {
	fx.New(
		Inject(),
		Invoke(),
	).Run()
}

// where invoke is always running in background
func Invoke() fx.Option {
	return fx.Invoke(
		StartHTTP,
		// database.InitDatabase,
	)
}

// injecting our file, like usecase, repository, routes, and etc
func Inject() fx.Option {
	return fx.Options(
		fx.Provide(StartHTTP, database.InitDatabase),
		usecase.Module,
		repository.Module,
	)
}

func StartHTTP(lc fx.Lifecycle, params Params) *http.Server {
	//initiate address server
	srv := &http.Server{Addr: ":8080"}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			StartServer(params, srv)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			StopServer(ctx, srv)
			return nil
		},
	})
	return srv
}

func StartServer(params Params, srv *http.Server) {
	ginRouter := gin.Default()

	// add middleware
	ginRouter.Use(middleware.RateLimiter())

	// for register router into server
	srv.Handler = RegisterRoutes(ginRouter, params)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func StopServer(ctx context.Context, srv *http.Server) {
	stop := NewStop(ctx, func() error {
		// do something
		return nil
	}, srv)

	stop.Execute()
}
