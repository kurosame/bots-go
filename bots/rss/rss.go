package rss

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/slack-go/slack"
)

func FilterTwitterRSS(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	user := slack.New(os.Getenv("SLACK_USER_OAUTH_TOKEN"), slack.OptionDebug(true))
	bot := slack.New(os.Getenv("SLACK_BOT_OAUTH_TOKEN"), slack.OptionDebug(true))

	conversation, err := bot.GetConversationHistory(
		&slack.GetConversationHistoryParameters{
			ChannelID: os.Getenv("SLACK_CHANNEL_ID"),
			Oldest:    "1654231477.168869",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	client, err := datastore.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	keywords, err := client.GetAll(ctx, datastore.NewQuery("Keyword").Namespace("TwitterRSSFilter").KeysOnly(), nil)
	if err != nil {
		log.Fatal(err)
	}

	excludeRe := regexp.MustCompile("<https?://.+>")

	if len(conversation.Messages) > 0 {
		fmt.Printf("Latest timestamp: %s\n", conversation.Messages[0].Timestamp)
	}

	for _, m := range conversation.Messages {
		text := strings.ToLower(excludeRe.ReplaceAllString(m.Text, ""))

		for _, k := range keywords {
			if strings.Contains(text, k.Name) {
				user.AddStar(os.Getenv("SLACK_CHANNEL_ID"), slack.ItemRef{Timestamp: m.Timestamp})
				time.Sleep(time.Second * 3)
				break
			}
		}
	}
}
