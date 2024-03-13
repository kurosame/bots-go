package rss

import "net/http"

func SetToken(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")

	print(token)
}
