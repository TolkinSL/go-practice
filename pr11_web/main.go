package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)

	log.Println("Server start on :8080")
	http.ListenAndServe(":8080", nil)
}

//Request - запрос клиента
//Response - ответ сервера

func homePage(w http.ResponseWriter, r *http.Request) {
	html := `<strong>Hello Go!</strong>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}
