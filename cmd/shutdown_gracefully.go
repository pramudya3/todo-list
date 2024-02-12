package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-list-app/internal/utils"
)

type stop struct {
	wait   chan struct{}
	f      func() error
	ctx    context.Context
	server *http.Server
}

func NewStop(ctx context.Context, f func() error, server *http.Server) *stop {
	return &stop{
		wait:   make(chan struct{}, 1),
		f:      f,
		ctx:    ctx,
		server: server,
	}
}

func (s *stop) Execute() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		defer func() {
			// wait until get signal
			<-s.wait
		}()
		// running process func
		s.f()
	}()

	fmt.Println("Signal killing recieved ..")
	select {
	case <-s.wait:
		return
	case <-exit:
		ctxCancel, cancel := context.WithTimeout(s.ctx, time.Second*time.Duration(utils.TimeoutContext))
		defer cancel()

		if err := s.server.Shutdown(ctxCancel); err != nil {
			log.Fatal("Server Shutdown Error:", err)
		}

		select {
		case <-ctxCancel.Done():
			fmt.Println("Timeout exceed, force killing server ..")
			return
		case <-s.wait:
			fmt.Println("Shutdown gracefully ..")
		}
	}
}
