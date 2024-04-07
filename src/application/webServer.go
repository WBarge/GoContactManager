package application

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type WebServer struct {
	RestServerPort uint16
}

func NewWebServer(config Config) *WebServer {
	returnValue := &WebServer{
		RestServerPort: config.RestServerPort}
	return returnValue
}

func (w *WebServer) Start(ctx context.Context) error {

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", w.RestServerPort),
		Handler: w.RegisterHandlers(),
	}
	fmt.Println("Starting server")

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
