package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type handler struct{}

func NewHandler() Handler {
	return &handler{}
}

func (*handler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	go task(ctx)

	fmt.Fprint(w, "/test: response\n")
}

func task(ctx context.Context) {
	time.Sleep(1 * time.Second)

	select {
	case <-ctx.Done():
		log.Println("context done.")
	default:
		log.Println("context not done.")
	}

	log.Println("task done.")
}
