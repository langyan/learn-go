package main

import (
	"log"
	"net/http"

	"github.com/langyan/lin-go-reverse-proxy/proxy"
)

func main() {
	targets := []string{
		"http://localhost:8081",
		"http://localhost:8082",
	}
	p := proxy.NewProxy(targets)
	log.Println("🚀 Reverse proxy started on :8080")
	log.Fatal(http.ListenAndServe(":8080", p))
}
