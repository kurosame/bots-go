resource "google_project_service" "this" {
  for_each = toset([
    "cloudbuild.googleapis.com",
    "iam.googleapis.com",
    "secretmanager.googleapis.com"
  ])
  service                    = each.key
  disable_dependent_services = true
  disable_on_destroy         = true
}
