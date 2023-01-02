package main

import (
	"context"
	"log"
	"os"

	"github.com/Symthy/golang-rest-web-app-study/internal/server"
)

func main() {
	if err := server.Run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
