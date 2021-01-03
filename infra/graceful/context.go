package graceful

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// NewContext create context that consider graceful shutdown.
func NewContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		sig := <-sigint
		log.Printf("Sig %s => GracefulStop triggered", sig.String())
		cancel()
	}()

	return ctx
}
