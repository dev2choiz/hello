
## Jenkins
resource "google_service_account" "sa-jenkins" {
  account_id   = "sa-jenkins"
  display_name  = "Jenkins Service Account"
}
resource "google_project_iam_member" "sa-jenkins-st-adm" {
  role    = "roles/storage.admin"
  member = "serviceAccount:${google_service_account.sa-jenkins.email}"
}
resource "google_project_iam_member" "sa-jenkins-cb-builder" {
  role    = "roles/cloudbuild.builds.builder"
  member = "serviceAccount:${google_service_account.sa-jenkins.email}"
}
resource "google_service_account_key" "sa-jenkins-key" {
  service_account_id = google_service_account.sa-jenkins.name
}
resource "kubernetes_secret" "sa-jenkins-credentials" {
  metadata {
    name      = "sa-jenkins-credentials"
    namespace = var.namespace
  }
  data = {
    "credentials.json" = base64decode(google_service_account_key.sa-jenkins-key.private_key)
  }
}

## Api
resource "google_service_account" "sa-hello-api" {
  account_id   = "sa-hello-api"
  display_name = "Hello-api Service Account"
}
resource "google_project_iam_member" "sa-hello-api-ps-publisher" {
  role    = "roles/pubsub.publisher"
  member = "serviceAccount:${google_service_account.sa-hello-api.email}"
}
resource "google_service_account_key" "sa-hello-key" {
  service_account_id = google_service_account.sa-hello-api.name
}
resource "kubernetes_secret" "sa-hello-api-credentials" {
  metadata {
    name      = "sa-hello-api-credentials"
    namespace = var.namespace
  }
  data = {
    "credentials.json" = base64decode(google_service_account_key.sa-hello-key.private_key)
  }
}

## Serverless
resource "google_service_account" "sa_serverless" {
  account_id   = "sa-serverless"
  display_name = "Serverless Service Account"
}
