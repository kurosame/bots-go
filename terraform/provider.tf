terraform {
  required_version = "1.7.4"

  cloud {
    organization = "kurosame"

    workspaces {
      name = "bots-go"
    }
  }

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.18.0"
    }
  }
}

provider "google" {
  project = var.GOOGLE_PROJECT_ID
  region  = "asia-northeast1"
}
