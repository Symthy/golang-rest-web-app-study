package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/Symthy/golang-rest-web-app-study/internal/server"
)

const (
	defaultPort = "18080"
)

func main() {
	p := defaultPort
	if len(os.Args) >= 2 {
		p = os.Args[1]
	}
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("failed to lisen port %s: %v", p, err)
	}

	if err := server.Run(context.Background(), l); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
