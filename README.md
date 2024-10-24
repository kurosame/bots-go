# Bots with Go

## [Twitter RSS Filter](https://github.com/kurosame/bots-go/tree/main/bots/rss)

Filter the Twitter list for RSS

To run it, create `.env` file then set the following

```sh
GOOGLE_PROJECT_ID            # GCP project id
GOOGLE_PROJECT_NUMBER        # GCP project number
SLACK_USER_OAUTH_TOKEN       # Slack user oauth token
SLACK_BOT_OAUTH_TOKEN        # Slack bot oauth token
SLACK_CHANNEL_ID_RSS         # Slack channel id (for #event-rss)
SLACK_CHANNEL_ID_TWITTER     # Slack channel id (for #media-twitter)
SLACK_CHANNEL_ID_TWITTER_RSS # Slack channel id (for #media-twitter-rss)
RSSAPP_ID_TWITTER            # RSS.app id (for media-twitter)
RSSAPP_ID_LIKE               # RSS.app id (for news-like)
```

Deploy to the Cloud Functions are as follows

```sh
make zip
make apply
```

### Run manually

```sh
# Add filter keyword
curl -X POST \
-H "Authorization: Bearer $(gcloud auth print-identity-token)" \
-H "Content-Type: application/json" \
https://asia-northeast1-【GCP ProjectID】.cloudfunctions.net/twitter-rss-filter-add-keyword\?kw\=kw1,kw2,kw3

# Set RSS.app token
curl -X POST \
-H "Authorization: Bearer $(gcloud auth print-identity-token)" \
-H "Content-Type: application/json" \
https://asia-northeast1-【GCP ProjectID】.cloudfunctions.net/twitter-rss-filter-set-token\?token\=token_value
```

## [JSON to Firestore](https://github.com/kurosame/json2firestore)

Read jsonl file and add/update to Firestore  
Not a generic implementation

## Terraform

### IAM setup for Terraform

Created a `bots-go-tf` IAM service account from the GCP Console  
`bots-go-tf` has a owner roles of the GCP resources to attached

Create a service account key and set `GOOGLE_CREDENTIALS` in Terraform Cloud Variables  
credential json needs to be changed as follows

```sh
cat credential.json | tr -s '\n' ' ' # Remove line feed code
```

### Terraform is not supported

The following resources are not supported by Terraform  
So create it from the GCP Console

- Create a Datastore entity
