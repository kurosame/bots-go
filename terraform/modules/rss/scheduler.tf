resource "google_cloud_scheduler_job" "twitter_rss_filter" {
  name     = "twitter-rss-filter"
  schedule = "00 * * * *"

  retry_config {
    retry_count = 3
  }

  http_target {
    http_method = "POST"
    uri         = google_cloudfunctions_function.twitter_rss_filter.https_trigger_url
    headers = {
      "Content-Type" = "application/json"
    }
    oidc_token {
      service_account_email = google_service_account.this.email
    }
  }
}

resource "google_cloud_scheduler_job" "twitter_rss_filter_clear_cache" {
  name     = "twitter-rss-filter-clear-cache"
  schedule = "*/15 * * * *"

  retry_config {
    retry_count = 3
  }

  http_target {
    http_method = "POST"
    uri         = google_cloudfunctions_function.twitter_rss_filter_clear_cache.https_trigger_url
    headers = {
      "Content-Type" = "application/json"
    }
    oidc_token {
      service_account_email = google_service_account.this.email
    }
  }
}
