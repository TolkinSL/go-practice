package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Go!")
}

func main() {
	fmt.Println("Server start!")
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8080", nil)
}
