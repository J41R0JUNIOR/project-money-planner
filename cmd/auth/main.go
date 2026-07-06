package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/J41R0JUNIOR/project-money-planner/internal/auth"
	"github.com/J41R0JUNIOR/project-money-planner/internal/bootstrap"
)

func main() {
	app, err := bootstrap.New(context.Background())
	if err != nil {
		panic(err)
	}

	service := auth.NewService()
	_ = service

	mux := http.NewServeMux()
	mux.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"message":"auth service ready"}`)
	})

	fmt.Println("auth service listening on :8083")
	log.Fatal(http.ListenAndServe(":8083", mux))
	_ = app
}
