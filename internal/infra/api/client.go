package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetCepFromApi(url string, ch chan map[string]interface{}) {
	fmt.Println("Requesting", url)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		panic(err)
	}

	jsonResponse["api"] = url
	fmt.Println("Response from", url)

	ch <- jsonResponse
}
