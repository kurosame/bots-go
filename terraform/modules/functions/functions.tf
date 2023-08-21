resource "google_cloudfunctions_function" "twitter_rss_filter" {
  name                  = "twitter-rss-filter"
  runtime               = "go116"
  available_memory_mb   = 128
  timeout               = 540
  source_archive_bucket = google_storage_bucket.this.name
  source_archive_object = google_storage_bucket_object.this.name
  trigger_http          = true
  entry_point           = "FilterTwitterRSS"

  environment_variables = {
    GCP_PROJECT_ID         = var.GOOGLE_PROJECT_ID
    SLACK_USER_OAUTH_TOKEN = var.SLACK_USER_OAUTH_TOKEN
    SLACK_BOT_OAUTH_TOKEN  = var.SLACK_BOT_OAUTH_TOKEN
    SLACK_CHANNEL_ID       = var.SLACK_CHANNEL_ID
    SLACK_CHANNEL_ID_RSS   = var.SLACK_CHANNEL_ID_RSS
  }
}

resource "google_cloudfunctions_function_iam_member" "twitter_rss_filter" {
  project        = google_cloudfunctions_function.twitter_rss_filter.project
  region         = google_cloudfunctions_function.twitter_rss_filter.region
  cloud_function = google_cloudfunctions_function.twitter_rss_filter.name
  role           = "roles/cloudfunctions.invoker"
  member         = "serviceAccount:${google_service_account.this.email}"
}

resource "google_cloudfunctions_function" "twitter_rss_filter_add_keyword" {
  name                  = "twitter-rss-filter-add-keyword"
  runtime               = "go116"
  available_memory_mb   = 128
  timeout               = 60
  source_archive_bucket = google_storage_bucket.this.name
  source_archive_object = google_storage_bucket_object.this.name
  trigger_http          = true
  entry_point           = "AddKeyword"

  environment_variables = {
    GCP_PROJECT_ID = var.GOOGLE_PROJECT_ID
  }
}

resource "google_cloudfunctions_function_iam_member" "twitter_rss_filter_add_keyword" {
  project        = google_cloudfunctions_function.twitter_rss_filter_add_keyword.project
  region         = google_cloudfunctions_function.twitter_rss_filter_add_keyword.region
  cloud_function = google_cloudfunctions_function.twitter_rss_filter_add_keyword.name
  role           = "roles/cloudfunctions.invoker"
  member         = "serviceAccount:${google_service_account.this.email}"
}
