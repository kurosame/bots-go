package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kurosame/bots-go/bots/rss"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", rss.FilterTwitterRSS)
	http.HandleFunc("/kw", rss.AddKeyword)
	http.HandleFunc("/token", rss.SetToken)
	http.HandleFunc("/cache", rss.ClearCache)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
