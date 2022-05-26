terraform {
  backend "remote" {
    organization = "kurosame"

    workspaces {
      name = "bots-go"
    }
  }
}

provider "google" {
  project = var.GOOGLE_PROJECT_ID
  region  = "asia-northeast1"
}
