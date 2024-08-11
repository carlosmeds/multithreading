package webserver

import (
	"fmt"
	"net/http"
	"time"

	"encoding/json"

	"github.com/carlosmeds/multithreading/internal/infra/api"
	"github.com/go-chi/chi/v5"
)

const (
	viaCep    = "http://viacep.com.br/ws/%s/json/"
	brasilApi = "https://brasilapi.com.br/api/cep/v1/%s"
)

type Error struct {
	Message string `json:"message"`
}

func GetCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")

	if len(cep) != 8 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "Invalid CEP"})
		return
	}

	ch1 := make(chan map[string]interface{})
	ch2 := make(chan map[string]interface{})
	viaCepUrl := fmt.Sprintf(viaCep, cep)
	brasilApiUrl := fmt.Sprintf(brasilApi, cep)

	go api.GetCepFromApi(viaCepUrl, ch1)
	go api.GetCepFromApi(brasilApiUrl, ch2)

	select {
	case result := <-ch1:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(result)
		return
	case result := <-ch2:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(result)
		return

	case <-time.After(time.Second * 1):
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusRequestTimeout)
		json.NewEncoder(w).Encode(Error{Message: "API timeout"})
		return
	}
}
