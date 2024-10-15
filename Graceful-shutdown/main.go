package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	s := NewServer(9000)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go s.Start(wg, ctx)

	wg.Wait()
}

type Server struct {
	listenAddr string
	server     *http.Server
}

func NewServer(addr int) *Server {
	server := &http.Server{Addr: fmt.Sprintf(":%d", addr)}

	return &Server{
		listenAddr: fmt.Sprintf(":%d", addr),
		server:     server,
	}
}

func (s *Server) Start(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()


	go func() {
		fmt.Println("Server Running!!")
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe Error: %v\n", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("Shutting down server gracefully...")


	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Server shutdown error: %v\n", err)
	}

	fmt.Println("Server shutdown complete!")
}
