package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"passwordGenerator/handlers"
	"passwordGenerator/internal/middleware"
	"syscall"
	"time"
)

func main() {

	router := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.LoggingMiddleware(router),
	}

	// Requests
	router.HandleFunc("/", handlers.PasswordHandler)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Server starting on port 8080...")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server exited gracefully.")
}
