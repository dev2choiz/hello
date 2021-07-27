provider "google" {
  project = var.project_id
  region  = var.region
  zone    = var.zone
}

provider "google-beta" {
  project = var.project_id
  region  = var.region
  zone    = var.zone
}

data "google_client_config" "provider" {}

provider "kubernetes" {
  host  = "https://${google_container_cluster.primary.endpoint}"
  token = data.google_client_config.provider.access_token
  cluster_ca_certificate = base64decode(google_container_cluster.primary.master_auth[0].cluster_ca_certificate)
}

# Enable Gcp services
resource "google_project_service" "project-enable-container" {
  project = var.project_id
  disable_on_destroy = false
  service = "container.googleapis.com"
}
resource "google_project_service" "project-enable-containeranalysis" {
  project = var.project_id
  disable_on_destroy = false
  service = "containeranalysis.googleapis.com"
}
resource "google_project_service" "project-enable-cloudfunctions" {
  project = var.project_id
  disable_on_destroy = false
  service = "cloudfunctions.googleapis.com"
}
resource "google_project_service" "project-enable-pubsub" {
  project = var.project_id
  disable_on_destroy = false
  service = "pubsub.googleapis.com"
}
resource "google_project_service" "project-enable-cloudkms" {
  project = var.project_id
  disable_on_destroy = false
  service = "cloudkms.googleapis.com"
}
resource "google_project_service" "project-enable-cloudbuild" {
  project = var.project_id
  disable_on_destroy = false
  service = "cloudbuild.googleapis.com"
}
resource "google_project_service" "project-enable-dns" {
  project = var.project_id
  disable_on_destroy = false
  service = "dns.googleapis.com"
}
resource "google_project_service" "project-enable-domains" {
  project = var.project_id
  disable_on_destroy = false
  service = "domains.googleapis.com"
}
resource "google_project_service" "project-enable-endpoints" {
  project = var.project_id
  disable_on_destroy = false
  service = "endpoints.googleapis.com"
}

resource "google_service_account" "cluster_nodes" {
  account_id   = "sa-cluster-nodes"
  display_name = "Cluster Nodes Service Account"
}
resource "google_project_iam_member" "sa-jenkins-cb-builder" {
  role    = "roles/editor"
  member = "serviceAccount:${google_service_account.cluster_nodes.email}"
}

# GKE cluster
resource "google_container_cluster" "primary" {
  name     = var.cluster_name
  location = var.zone

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1

  network    = google_compute_network.vpc.name
  subnetwork = google_compute_subnetwork.subnet.name
  depends_on = [google_project_service.project-enable-container]
}

# Separately Managed Node Pool
resource "google_container_node_pool" "primary_nodes" {
  name       = "${google_container_cluster.primary.name}-node-pool"
  location   = var.zone
  cluster    = google_container_cluster.primary.name
  node_count = var.gke_num_nodes

  node_config {
    preemptible  = true
    machine_type = var.machine_type

    service_account = google_service_account.cluster_nodes.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform",
    ]

    labels = {
      env = var.project_id
    }

    disk_size_gb = var.disk_size
    tags         = ["gke-node", "${var.project_id}-gke"]
    metadata = {
      disable-legacy-endpoints = "true"
    }
  }
}

resource "kubernetes_namespace" "cluster_namespace" {
  metadata {
    name = var.namespace
  }
}

# DNS
resource "google_dns_managed_zone" "main-zone" {
  provider    = google-beta
  name        = "main-zone"
  dns_name    = "${var.domain}."
  description = "Main DNS zone"
  labels = {}
  visibility = "public"

  dnssec_config {
    state = "off"
  }
}
