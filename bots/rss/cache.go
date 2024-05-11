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
	"github.com/slack-go/slack"
)

func ClearCache(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	bot := slack.New(os.Getenv("SLACK_BOT_OAUTH_TOKEN"), slack.OptionDebug(true))

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
	var textTwitter slack.MsgOption
	if err == nil {
		textTwitter = slack.MsgOptionText(fmt.Sprintf("Twitter response: %s", string(resTwitter)), false)
	} else {
		textTwitter = slack.MsgOptionText(fmt.Sprintf("Twitter response error: %s", err.Error()), false)
	}
	if _, _, err := bot.PostMessage(os.Getenv("SLACK_CHANNEL_ID_RSS"), textTwitter, slack.MsgOptionAsUser(false)); err != nil {
		log.Fatal(err)
	}

	resLike, err := clientGql.MutateRaw(ctx, m, varLike, graphql.OperationName("clearFeedCache"))
	var textLike slack.MsgOption
	if err == nil {
		textLike = slack.MsgOptionText(fmt.Sprintf("Like response: %s", string(resLike)), false)
	} else {
		textLike = slack.MsgOptionText(fmt.Sprintf("Like response error: %s", err.Error()), false)
	}
	if _, _, err := bot.PostMessage(os.Getenv("SLACK_CHANNEL_ID_RSS"), textLike, slack.MsgOptionAsUser(false)); err != nil {
		log.Fatal(err)
	}
}
