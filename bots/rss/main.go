package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	user := slack.New(os.Getenv("SLACK_USER_OAUTH_TOKEN"), slack.OptionDebug(true))
	bot := slack.New(os.Getenv("SLACK_BOT_OAUTH_TOKEN"), slack.OptionDebug(true))

	conversation, err := bot.GetConversationHistory(
		&slack.GetConversationHistoryParameters{
			ChannelID: os.Getenv("SLACK_CHANNEL_ID"),
			Oldest:    "1653234276.300029",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	keywords := []string{
		"cloudflare", "javascript", "go", "html", "python", "figma", "値オブジェクト", "react", "デザイン",
		"remix", "worker", "tailwind", "eslint", "prettier", "vitest", "node",
		"d1", "r2", "kv", "swift", "function", "ios", "db", "cdn", "ddd", "aws", "vite",
		"postgre", "sql", "cloud", "action", "ci", "cd", "web", "www", "edge", "rust",
	}
	excludeRe := regexp.MustCompile("<https?://.+>")

	if len(conversation.Messages) > 0 {
		fmt.Printf("Latest timestamp: %s\n", conversation.Messages[0].Timestamp)
	}

	for _, m := range conversation.Messages {
		text := strings.ToLower(excludeRe.ReplaceAllString(m.Text, ""))

		for _, k := range keywords {
			if strings.Contains(text, k) {
				user.AddStar(os.Getenv("SLACK_CHANNEL_ID"), slack.ItemRef{Timestamp: m.Timestamp})
				time.Sleep(time.Second * 3)
				break
			}
		}
	}
}
