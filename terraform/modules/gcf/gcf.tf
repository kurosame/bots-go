resource "google_cloudfunctions_function" "this" {
  name                  = "twitter-rss-filter"
  runtime               = "go116"
  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.this.name
  source_archive_object = google_storage_bucket_object.this.name
  trigger_http          = true
  entry_point           = "FilterTwitterRSS"
}

resource "google_cloudfunctions_function_iam_member" "this" {
  project        = google_cloudfunctions_function.this.project
  region         = google_cloudfunctions_function.this.region
  cloud_function = google_cloudfunctions_function.this.name
  role           = "roles/cloudfunctions.invoker"
  member         = "serviceAccount:${google_service_account.this.email}"
}
