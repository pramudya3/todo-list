package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-list-app/database"
	"todo-list-app/domain"
	"todo-list-app/http/router"
	"todo-list-app/repository"
	"todo-list-app/usecase"

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
	srv := &http.Server{Addr: ":8080"}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			StartServer(params, srv)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			StopServer(srv)
			return nil
		},
	})
	return srv
}

func StartServer(params Params, srv *http.Server) {
	// start server gracefully
	ginServer := gin.Default()

	// initiate routers, if there is new
	// router's group, just add in this line
	router.NewHealthcheckRoutes(ginServer.Group("/healthcheck"))
	router.NewUserRoutes(ginServer.Group("/users"), params.UserUsecase)

	// for register router into server
	srv.Handler = ginServer
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

func StopServer(srv *http.Server) {
	// Stop server gracefully

	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
