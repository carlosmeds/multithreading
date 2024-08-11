package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	viaCepUrl    = "https://viacep.com.br/ws/01001000/json/"
	brasilApiUrl = "https://brasilapi.com.br/api/cep/v1/01001000"
)

func getCep(url string, ch chan string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	ch <- string(body)

}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go getCep(viaCepUrl, ch1)
	go getCep(brasilApiUrl, ch2)

	select {
	case <-ch1:
		println("ViaCep")
	case <-ch2:
		println("BrasilApi")
	case <-time.After(time.Second * 1):
		println("Timeout")
	}
}
