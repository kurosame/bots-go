package rss

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/slack-go/slack"
)

func FilterTwitterRSS(w http.ResponseWriter, r *http.Request) {
	user := slack.New(os.Getenv("SLACK_USER_OAUTH_TOKEN"), slack.OptionDebug(true))
	bot := slack.New(os.Getenv("SLACK_BOT_OAUTH_TOKEN"), slack.OptionDebug(true))

	conversation, err := bot.GetConversationHistory(
		&slack.GetConversationHistoryParameters{
			ChannelID: os.Getenv("SLACK_CHANNEL_ID"),
			Oldest:    "1653807277.376809",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	keywords := []string{
		"cloudflare", "javascript", "go", "html", "python", "figma", "値オブジェクト", "react", "デザイン",
		"remix", "worker", "tailwind", "eslint", "prettier", "vitest", "node", "alb",
		"d1", "r2", "kv", "swift", "function", "ios", "db", "cdn", "ddd", "aws", "vite",
		"postgre", "sql", "cloud", "action", "ci", "cd", "web", "www", "edge", "rust",
		"github", "terraform", "redux", "typescript", "next", "v8", "deno", "firebase",
		"webrtc", "分散", "コンテキスト", "css", "ux", "ブラウザ", "glue", "etl",
		"tls", "component", "vue", "spa", "vscode", "firefox", "monorepo", "composition",
		"bigquery", "フロント", "サーバ", "インフラ", "gcp", "orm", "design", "angular", "snapshot",
		"puppeteer", "ui", "os", "recoil", "router", "i18n", "dart", "chromium", "m1", "quic",
		"layout", "storybook", "pip", "hook", "domain", "vuex", "flutter", "graphql", "nuxt",
		"mac", "kubernetes", "k8s", "di", "android", "npm", "ec2", "netlify", "コンテナ", "ログ",
		"docker", "webgl", "opengl", "lambda", "gateway", "nginx", "esm", "babel", "esbuild",
		"webpack", "dom", "iframe", "scala", "apache", "s3", "gcs", "interface", "type", "xss",
		"a11y", "atomic", "分散", "svelte", "パフォーマンス", "redis", "swr", "webassembly", "js", "ts",
		"テスト", "ssr", "usecase", "ユースケース", "firestore", "parcel", "cloudfront", "hmr", "map",
		"jsx", "mongo", "memory", "メモリ", "spark", "gatsby", "amplify", "cloudwatch", "auth",
		"cypress", "carthage", "ブロックチェーン", "blockchain", "aurora", "fiber",
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
