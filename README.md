# Bots with Go

## [Twitter RSS Filter](https://github.com/kurosame/bots-go/tree/main/bots/rss)

Filter the Twitter list for RSS

To run it, create `.env` file then set the following

```sh
GCP_PROJECT_ID         # GCP project id
SLACK_USER_OAUTH_TOKEN # Slack user oauth token
SLACK_BOT_OAUTH_TOKEN  # Slack bot oauth token
SLACK_CHANNEL_ID       # Slack channel to filter
```

## IAM setup for Terraform

Created a `bots-go-tf` IAM service account from the GCP Console  
`bots-go-tf` has a owner roles of the GCP resources to attached

Create a service account key and set `GOOGLE_CREDENTIALS` in Terraform Cloud Variables  
credential json needs to be changed as follows

```sh
cat credential.json | tr -s '\n' ' ' # Remove line feed code
```

## Terraform is not supported

The following resources are not supported by Terraform  
So create it from the GCP Console

- Create a Datastore entity
