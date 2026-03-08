package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbix"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snppet"))
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		// Set a new cache-control header. If an existing Cache-Control header exists 
// it will be overwritten. 
w.Header().Set( Cache-Control, public, max-age=31536000) 
// In contrast, the Add() method appends a new Cache-Control header and can 
// be called multiple times. 
w.Header().Add( Cache-Control, public) 
w.Header().Add( Cache-Control, max-age=31536000) 
// Delete all values for the Cache-Control header. 
w.Header().Del( Cache-Control) 
// Retrieve the first value for the Cache-Control header. 
w.Header().Get( Cache-Control)
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on: 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
