terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 3.88.0"
    }
    google-beta = {
      source = "hashicorp/google-beta"
      version = "~> 3.88.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.5.1"
    }
  }

  backend "gcs" {
    bucket  = "d2c-tf-state"
    prefix  = "hello/hello-api"
  }

  required_version = "~> 1.0.3"
}
