package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
)

func main() {
	handler := cors.AllowAll().Handler(routes())

	srv := &http.Server{
		Addr:         ":1337",
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Listening on %s...", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
