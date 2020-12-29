package main

import (
	"net/http"
	handler "vedioAuth/handler"
)

func main() {
	http.HandleFunc("/", handler.WithArgHandler)
	http.ListenAndServe(":8000", nil)
}
