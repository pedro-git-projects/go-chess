package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
	})
	handler := c.Handler(routes())

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
