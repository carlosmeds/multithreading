package main

import (
	"fmt"
	"net/http"

	"github.com/carlosmeds/multithreading/internal/infra/webserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Server running on port 8000")
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ceps/{cep}", webserver.GetCep)

	http.ListenAndServe(":8000", r)
}
