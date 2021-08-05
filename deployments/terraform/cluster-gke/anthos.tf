# https://cloud.google.com/architecture/provisioning-anthos-clusters-with-terraform
module "project-services-anthos" {
  source  = "terraform-google-modules/project-factory/google//modules/project_services"

  project_id  = data.google_client_config.provider.project
  disable_services_on_destroy = false
  activate_apis = [
    "compute.googleapis.com",
    "iam.googleapis.com",
    "cloudresourcemanager.googleapis.com",
    "cloudtrace.googleapis.com",
    "meshca.googleapis.com",
    "meshtelemetry.googleapis.com",
    "meshconfig.googleapis.com",
    "iamcredentials.googleapis.com",
    "gkeconnect.googleapis.com",
    "gkehub.googleapis.com",
    "monitoring.googleapis.com",
    "logging.googleapis.com",
    "stackdriver.googleapis.com",
    //"anthos.googleapis.com",
  ]
}

module "hub-primary" {
  source           = "terraform-google-modules/kubernetes-engine/google//modules/hub"

  project_id       = data.google_client_config.provider.project
  cluster_name     = google_container_cluster.primary.name
  location         = google_container_cluster.primary.location
  cluster_endpoint = google_container_cluster.primary.endpoint
  gke_hub_membership_name = "primary"
  gke_hub_sa_name         = "primary"
}

module "asm-primary" {
  source           = "terraform-google-modules/kubernetes-engine/google//modules/asm"
  version          = "16.0.1"
  project_id       = data.google_client_config.provider.project
  cluster_name     = google_container_cluster.primary.name
  location         = google_container_cluster.primary.location
  cluster_endpoint = google_container_cluster.primary.endpoint
  enable_cluster_roles  = true
  enable_gcp_apis       = false

  outdir           = "asm-dir-${google_container_cluster.primary.name}"
  depends_on = [module.project-services-anthos]
}
