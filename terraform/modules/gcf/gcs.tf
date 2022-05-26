resource "google_storage_bucket" "this" {
  name          = "twitter-rss-filter"
  location      = "ASIA-NORTHEAST1"
  storage_class = "STANDARD"
}

resource "google_storage_bucket_object" "this" {
  name   = "rss.zip"
  bucket = google_storage_bucket.this.name
  source = "${path.module}/rss.zip"
}
