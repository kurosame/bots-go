resource "google_service_account" "this" {
  account_id = "rss-invoker"
}

resource "google_project_iam_member" "this" {
  for_each = toset([
    "roles/datastore.owner",
    "roles/secretmanager.secretAccessor",
    "roles/secretmanager.secretVersionAdder"
  ])

  project = var.GOOGLE_PROJECT_ID
  role    = each.key
  member  = "serviceAccount:${google_service_account.this.email}"
}
