package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	// _ "project_auth_jwt/docs"

	"project_auth_jwt/infra"
	"project_auth_jwt/routes"
	"syscall"
	"time"
)

func main() {
	// testing.MainTesting()
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	// init cron
	// cmd.CronJob(&ctx.Ctl)

	r := routes.NewRoutes(*ctx)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// Start the server
		log.Printf("Server running on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// Create a timeout context for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// Catching context timeout
	select {
	case <-shutdownCtx.Done():
		log.Println("Timeout of 5 seconds.")
	}

	log.Println("Server exiting")
}
