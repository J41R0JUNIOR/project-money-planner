package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/J41R0JUNIOR/project-money-planner/internal/bootstrap"
	userpkg "github.com/J41R0JUNIOR/project-money-planner/internal/modules/user"
)

func main() {
	app, err := bootstrap.New(context.Background())
	if err != nil {
		panic(err)
	}

	service := userpkg.NewService()
	handler := userpkg.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", handler.Create)
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetByID(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	fmt.Println("user service listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
	_ = app
}
