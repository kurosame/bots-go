resource "google_cloud_scheduler_job" "this" {
  name     = "twitter-rss-filter"
  schedule = "00 * * * *"

  retry_config {
    retry_count = 3
  }

  http_target {
    http_method = "POST"
    uri         = google_cloudfunctions_function.this.https_trigger_url
    headers = {
      "Content-Type" = "application/json"
    }
    oidc_token {
      service_account_email = google_service_account.this.email
    }
  }
}
