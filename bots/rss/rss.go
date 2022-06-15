package rss

import (
	"context"
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

type Timestamp struct {
	Timestamp string
}

func FilterTwitterRSS(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	user := slack.New(os.Getenv("SLACK_USER_OAUTH_TOKEN"), slack.OptionDebug(true))
	bot := slack.New(os.Getenv("SLACK_BOT_OAUTH_TOKEN"), slack.OptionDebug(true))

	client, err := datastore.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var tss []*Timestamp
	tsKeys, err := client.GetAll(ctx, datastore.NewQuery("Timestamp").Namespace("TwitterRSSFilter"), &tss)
	if err != nil {
		log.Fatal(err)
	}
	if len(tss) != 1 || len(tsKeys) != 1 {
		log.Fatal("Timestamp in Datastore is invalid")
	}

	conversation, err := bot.GetConversationHistory(
		&slack.GetConversationHistoryParameters{
			ChannelID: os.Getenv("SLACK_CHANNEL_ID"),
			Oldest:    tss[0].Timestamp,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	keywords, err := client.GetAll(ctx, datastore.NewQuery("Keyword").Namespace("TwitterRSSFilter").KeysOnly(), nil)
	if err != nil {
		log.Fatal(err)
	}

	excludeRe := regexp.MustCompile("<https?://.+>")

	if len(conversation.Messages) > 0 {
		if _, err := client.Put(ctx, tsKeys[0], &Timestamp{Timestamp: conversation.Messages[0].Timestamp}); err != nil {
			log.Fatal(err)
		}
	}

	for _, m := range conversation.Messages {
		text := strings.ToLower(excludeRe.ReplaceAllString(m.Text, ""))

		for _, k := range keywords {
			if strings.Contains(text, k.Name) {
				if err := user.AddStar(os.Getenv("SLACK_CHANNEL_ID"), slack.ItemRef{Timestamp: m.Timestamp}); err != nil {
					log.Fatal(err)
				}
				time.Sleep(time.Second * 3)
				break
			}
		}
	}
}
