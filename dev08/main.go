package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Hello Go!"))
	if err != nil {
		http.Error(w, "Ooops", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server start http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error in server start", err)
	}
}
