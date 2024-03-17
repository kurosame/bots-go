resource "google_secret_manager_secret" "this" {
  secret_id = "twitter-rss-filter"

  replication {
    auto {}
  }
}
