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
      version = "5.44.2"
    }
  }
}

provider "google" {
  project = var.GOOGLE_PROJECT_ID
  region  = "asia-northeast1"
}
