
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
## save google_service_account_key.sa-jenkins-key locally
resource "local_file" "file-sa-jenkins-key" {
    sensitive_content = base64decode(google_service_account_key.sa-jenkins-key.private_key)
    filename = "${var.dir_output}/sa-jenkins-credentials.json"
}

## Api
resource "google_service_account" "sa-hello-api" {
  account_id   = "sa-hello-api"
  display_name = "Hello-api Service Account"
}
resource "google_project_iam_member" "sa-hello-api-ps-publisher" {
  role   = "roles/pubsub.publisher"
  member = "serviceAccount:${google_service_account.sa-hello-api.email}"
}
resource "google_project_iam_member" "sa-hello-api-svc-ctrl" {
  role    = "roles/servicemanagement.serviceController"
  member = "serviceAccount:${google_service_account.sa-hello-api.email}"
}
resource "google_project_iam_member" "sa-hello-api-cl-ag" {
  role    = "roles/cloudtrace.agent"
  member = "serviceAccount:${google_service_account.sa-hello-api.email}"
}
resource "google_project_iam_member" "sa-hello-api-log-lw" {
  role    = "roles/logging.logWriter"
  member = "serviceAccount:${google_service_account.sa-hello-api.email}"
}
resource "google_project_iam_member" "sa-hello-api-sql-client" {
  role    = "roles/cloudsql.client"
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
## save google_service_account_key.sa-hello-key locally
resource "local_file" "file-sa-hello-key" {
    sensitive_content = base64decode(google_service_account_key.sa-hello-key.private_key)
    filename = "${var.dir_output}/sa-hello-api-credentials.json"
}

## Serverless
resource "google_service_account" "sa_serverless" {
  account_id   = "sa-serverless"
  display_name = "Serverless Service Account"
}
