package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"Content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Homepage")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests () {
	http.HandleFunc("/",homePage)
	http.HandleFunc("/articles",returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000",nil))
}

func main () {
	Articles = []Article{
		Article{Title: "Hello 1", Desc: "Article Desc", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Desc", Content: "Article Content"},
	}
	handleRequests()
}