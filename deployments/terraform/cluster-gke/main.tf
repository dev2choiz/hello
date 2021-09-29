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
module "project-services" {
  source  = "terraform-google-modules/project-factory/google//modules/project_services"

  project_id  = data.google_client_config.provider.project
  disable_services_on_destroy = false
  activate_apis = [
    "container.googleapis.com",
    "containeranalysis.googleapis.com",
    "cloudfunctions.googleapis.com",
    "pubsub.googleapis.com",
    "cloudkms.googleapis.com",
    "cloudbuild.googleapis.com",
    "dns.googleapis.com",
    "domains.googleapis.com",
    "endpoints.googleapis.com",
    "run.googleapis.com",
    "sql-component.googleapis.com",
  ]
}

resource "google_service_account" "cluster_nodes" {
  account_id   = "sa-cluster-nodes"
  display_name = "Cluster Nodes Service Account"
}
resource "google_project_iam_member" "sa-jenkins-cb-builder" {
  role    = "roles/editor"
  member = "serviceAccount:${google_service_account.cluster_nodes.email}"
}

data "google_project" "project" {}

# GKE cluster
resource "google_container_cluster" "primary" {
  provider = google-beta
  name     = var.cluster_name
  location = var.zone

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1

  network    = google_compute_network.vpc.name
  subnetwork = google_compute_subnetwork.subnet.name

  depends_on = [module.project-services]

  # anthos
  /*resource_labels = {
    mesh_id = "proj-${data.google_project.project.number}"
  }*/
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
    labels = {
      // anthos
      /*"istio.io/rev" = "196-2"
      istio-injection = "enabled"*/
    }
  }
}
