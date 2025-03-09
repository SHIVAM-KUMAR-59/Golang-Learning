package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

/*
URL Shortener API in Go

Description:
This is a simple URL shortener API built using Go. It provides functionalities to:
1. Shorten a given URL and store it in memory.
2. Retrieve the original URL using the shortened version.
3. Redirect users to the original URL when accessing the short URL.

Endpoints:
1. GET / -> Returns a welcome message.
2. POST /shorten -> Accepts a JSON payload with a URL and returns a shortened version.
3. GET /redirect/{shortURL} -> Redirects the user to the original URL.

Tech Stack:
- Go
- HTTP Handlers
- MD5 Hashing for URL Shortening
- In-Memory Storage (Map)
*/

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"originalURL"`
	ShortURL     string    `json:"shortURL"`
	CreationDate time.Time `json:"creation_date"`
}

// In-memory database to store shortened URLs
var URLDB = make(map[string]URL)

// generateShortURL creates a shortened version of the provided URL using MD5 hashing
func generateShortURL(originalURL string) (string, error) {
	if originalURL == "" || originalURL == " " {
		return "", errors.New("invalid URL")
	}

	hasher := md5.New()
	hasher.Write([]byte(originalURL)) // Converts the original URL string into a byte slice
	hash := hex.EncodeToString(hasher.Sum(nil)) // Converts the byte slice into a hexadecimal string

	return hash[:8], nil // Returns the first 8 characters of the hash as the short URL
}

// createURL generates and stores a short URL for the given original URL
func createURL(originalURL string) (string, error) {
	shortURL, err := generateShortURL(originalURL)
	if err != nil {
		return "", errors.New("invalid URL")
	}

	id := shortURL
	URLDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}

	return shortURL, nil
}

// getURL retrieves the original URL corresponding to the given short URL
func getURL(id string) (URL, error) {
	url, ok := URLDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

// handler returns a simple welcome message
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go URL Shortener!")
}

// ShortURLHandler handles the /shorten endpoint and generates a short URL
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Generate a short URL
	shortURL, createErr := createURL(data.URL)
	if createErr != nil {
		http.Error(w, createErr.Error(), http.StatusBadRequest)
		return
	}

	// Send the response
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// redirectURLHandler redirects users to the original URL based on the shortened URL
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):] // Extract the short URL from the request path

	url, getErr := getURL(id)
	if getErr != nil {
		http.Error(w, getErr.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	// Register the HTTP handler functions
	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	// Start the server
	fmt.Println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
