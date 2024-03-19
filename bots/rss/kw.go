package rss

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/datastore"
)

func contains(dks []*datastore.Key, s string) bool {
	for _, dk := range dks {
		if dk.Name == s {
			return true
		}
	}
	return false
}

func AddKeyword(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, os.Getenv("GOOGLE_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	kws, err := client.GetAll(ctx, datastore.NewQuery("Keyword").Namespace("TwitterRSSFilter").KeysOnly(), nil)
	if err != nil {
		log.Fatal(err)
	}

	qKws := strings.Split(r.URL.Query().Get("kw"), ",")
	var addKws []string
	for _, qk := range qKws {
		s := strings.TrimSpace(strings.ToLower(qk))
		if s != "" && !contains(kws, s) {
			key := datastore.NameKey("Keyword", s, nil)
			key.Namespace = "TwitterRSSFilter"
			if _, err := client.Put(ctx, key, &struct{}{}); err != nil {
				log.Fatal(err)
			}
			addKws = append(addKws, s)
		}
	}
	if err := json.NewEncoder(w).Encode(addKws); err != nil {
		log.Fatal(err)
	}
}
