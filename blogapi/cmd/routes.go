package main

import (
	"net/http"
)

func routes() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/posts", getPosts)
	m.HandleFunc("/post", getPost)
	return m
}
