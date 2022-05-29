package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kurosame/bots-go/bots/rss"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", rss.FilterTwitterRSS)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
