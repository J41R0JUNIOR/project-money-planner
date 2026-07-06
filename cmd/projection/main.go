package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/J41R0JUNIOR/project-money-planner/internal/bootstrap"
	projectionpkg "github.com/J41R0JUNIOR/project-money-planner/internal/modules/projection"
)

func main() {
	app, err := bootstrap.New(context.Background())
	if err != nil {
		panic(err)
	}

	service := projectionpkg.NewService()
	_ = service

	mux := http.NewServeMux()
	mux.HandleFunc("/projections", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprint(w, `{"message":"projection service ready"}`)
	})

	fmt.Println("projection service listening on :8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
	_ = app
}
