package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://pkg.go.dev/")
	if err != nil {
		log.Fatal("not retrieve site", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("not read body", err)
	}

	log.Println(string(body))
}
