package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	// Image []string `json:"images"`
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	posts := []Post{
		{ID: 1, Title: "First Post", Body: "Lorem ipsum dolor sit amet"},
		{ID: 2, Title: "Second Post", Body: "Consectetur adipiscing elit"},
		{ID: 3, Title: "Third Post", Body: "Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"},
	}
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	post := Post{ID: id, Title: "Sample Post", Body: "This is a sample post"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
