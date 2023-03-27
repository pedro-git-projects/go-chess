package main

import (
	"fmt"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}
	fmt.Println("Starting server on port 8080")
	srv.ListenAndServe()
}
