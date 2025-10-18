package main

import (
	"context"
	"hello-world-api/app/container"
	"hello-world-api/app/http/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx := context.Background()

	cfg := container.ResolveConfig()

	srv := server.Run(cfg)

	select {
	case <-sigs:
		server.Stop(ctx, srv)
	}
}
