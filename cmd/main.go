package main

import (
	"context"
	"log"
	"os"

	"github.com/Symthy/golang-rest-web-app-study/internal/server"
)

const (
	defaultPort = "18080"
)

func main() {
	// p := defaultPort
	// if len(os.Args) >= 2 {
	// 	p = os.Args[1]
	// }

	if err := server.Run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
