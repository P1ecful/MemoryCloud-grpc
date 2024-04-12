package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	wApp := fiber.New()

	//
	// start listening and graceful shutdown
	//
	go func() {
		if err := wApp.Listen(":1200"); err != nil {
			log.Fatal("Error while listening")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := wApp.Shutdown(); err != nil { // try to stop server
		log.Print("Failed to stop server")

		return
	}

	log.Print("Server stopped")
}
