resource "google_cloudfunctions_function" "this" {
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
  }
}

resource "google_cloudfunctions_function_iam_member" "this" {
  project        = google_cloudfunctions_function.this.project
  region         = google_cloudfunctions_function.this.region
  cloud_function = google_cloudfunctions_function.this.name
  role           = "roles/cloudfunctions.invoker"
  member         = "serviceAccount:${google_service_account.this.email}"
}
