package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/Symthy/golang-rest-web-app-study/internal/config"
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

	cfg, err := config.New()
	if err != nil {
		log.Printf("failed to configration: %v", err)
		os.Exit(1)
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Printf("failed to listen: %v", err)
		os.Exit(1)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})

	s := server.NewServer(l, mux)
	if err := s.Run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
