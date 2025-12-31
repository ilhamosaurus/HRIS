package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ilhamosaurus/HRIS/internal/app"
	"github.com/ilhamosaurus/HRIS/pkg/setting"
)

func main() {
	addr := fmt.Sprintf(":%d", setting.Server.Port)
	appInstance, err := app.NewApp(addr)
	if err != nil {
		log.Fatalf("failed to initiate hris app: %+v", err)
	}

	go func() {
		log.Printf("server started at: %d", setting.Server.Port)
		if err := appInstance.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := appInstance.Stop(ctx); err != nil {
		log.Fatalf("failed to shutdown server: %v", err)
	}
	log.Println("server stopped")
}
