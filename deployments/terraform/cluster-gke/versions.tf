terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 3.76.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.3.2"
    }
  }

  required_version = "~> 1.0.3"
}
