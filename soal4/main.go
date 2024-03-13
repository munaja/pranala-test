package main

// "github.com/labstack/echo/v4"

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/items", getItems)

	http.ListenAndServe(":8000", nil)
}
