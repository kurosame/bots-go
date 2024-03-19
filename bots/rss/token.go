package rss

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

func SetToken(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	req := &secretmanagerpb.AddSecretVersionRequest{
		Parent: fmt.Sprintf("projects/%s/secrets/twitter-rss-filter", os.Getenv("GOOGLE_PROJECT_NUMBER")),
		Payload: &secretmanagerpb.SecretPayload{
			Data: []byte(r.URL.Query().Get("token")),
		},
	}

	if _, err := client.AddSecretVersion(ctx, req); err != nil {
		log.Fatal(err)
	}
}
