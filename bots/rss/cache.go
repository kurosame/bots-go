package rss

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/hasura/go-graphql-client"
)

func ClearCache(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	clientSm, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer clientSm.Close()

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/twitter-rss-filter/versions/latest", os.Getenv("GOOGLE_PROJECT_NUMBER")),
	}

	res, err := clientSm.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	clientGql := graphql.NewClient("https://rss.app/gql", nil).WithRequestModifier(func(r *http.Request) {
		r.Header.Set("Content-Type", "application/json")
		r.Header.Set("Cookie", fmt.Sprintf("token=%s", string(res.Payload.Data)))
	})

	var m struct {
		ClearFeedCache graphql.ID `graphql:"clearFeedCache(id: $id)"`
	}
	varTwitter := map[string]interface{}{
		"id": graphql.ID(os.Getenv("RSSAPP_ID_TWITTER")),
	}
	varLike := map[string]interface{}{
		"id": graphql.ID(os.Getenv("RSSAPP_ID_LIKE")),
	}

	resTwitter, err := clientGql.MutateRaw(ctx, m, varTwitter, graphql.OperationName("clearFeedCache"))
	if err != nil {
		log.Fatal(err)
	}
	print(string(resTwitter))

	resLike, err := clientGql.MutateRaw(ctx, m, varLike, graphql.OperationName("clearFeedCache"))
	if err != nil {
		log.Fatal(err)
	}
	print(string(resLike))
}
