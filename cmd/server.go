package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"todo-list-app/database"
	"todo-list-app/repository"
	"todo-list-app/router"
	"todo-list-app/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

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
		fx.Provide(StartHTTP, router.NewRoutes, database.InitDatabase),
		usecase.Module,
		repository.Module,
	)
}

func StartHTTP(lc fx.Lifecycle, routes *gin.Engine) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: routes}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
