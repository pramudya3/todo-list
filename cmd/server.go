package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"todo-list-app/database"
	"todo-list-app/domain"
	middleware "todo-list-app/internal/http/midlleware"
	"todo-list-app/internal/repository"
	"todo-list-app/internal/usecase"
	"todo-list-app/internal/utils"

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
		fx.Provide(utils.LoadConfig, StartHTTP, database.InitDatabase),
		usecase.Module,
		repository.Module,
	)
}

func StartHTTP(lc fx.Lifecycle, params Params, cfg *domain.Config) *http.Server {
	//initiate address server
	srv := &http.Server{Addr: fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.ServerPort)}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			StartServer(params, srv)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func StartServer(params Params, srv *http.Server) {
	gin.SetMode(gin.ReleaseMode)
	ginRouter := gin.Default()

	// add middleware
	ginRouter.Use(middleware.RateLimiter())

	// for register router into server
	srv.Handler = RegisterRoutes(ginRouter, params)
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("error listen server: \n", err)
		}
	}()
	log.Println("server is listening at:", srv.Addr)
}

func StopServer(ctx context.Context, srv *http.Server) {
	stop := NewStop(ctx, func() error {
		// do something
		return nil
	}, srv)

	stop.Execute()
}
