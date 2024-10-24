module "rss" {
  source = "./modules/rss"

  GOOGLE_PROJECT_ID            = var.GOOGLE_PROJECT_ID
  GOOGLE_PROJECT_NUMBER        = var.GOOGLE_PROJECT_NUMBER
  SLACK_BOT_OAUTH_TOKEN        = var.SLACK_BOT_OAUTH_TOKEN
  SLACK_CHANNEL_ID_RSS         = var.SLACK_CHANNEL_ID_RSS
  SLACK_CHANNEL_ID_TWITTER     = var.SLACK_CHANNEL_ID_TWITTER
  SLACK_CHANNEL_ID_TWITTER_RSS = var.SLACK_CHANNEL_ID_TWITTER_RSS
  RSSAPP_ID_TWITTER            = var.RSSAPP_ID_TWITTER
  RSSAPP_ID_LIKE               = var.RSSAPP_ID_LIKE
}
