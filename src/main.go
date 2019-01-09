package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string
	Desc    string
	Content string
}

type Articles []Article

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	fmt.Fprintf(w, "Helsslo ddssworld")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}

func main() {
	handleRequests()
	fmt.Println("hello world")
}
