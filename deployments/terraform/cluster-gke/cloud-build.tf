data "google_project" "project" {}

resource "google_project_iam_member" "default-cb-iam-asu" {
  role    = "roles/iam.serviceAccountUser"
  member = "serviceAccount:${data.google_project.project.number}@cloudbuild.gserviceaccount.com"
  depends_on = [google_project_service.project-enable-cloudbuild]
}

resource "google_project_iam_member" "default-cb-cf-dev" {
  role    = "roles/cloudfunctions.developer"
  member = "serviceAccount:${data.google_project.project.number}@cloudbuild.gserviceaccount.com"
  depends_on = [google_project_service.project-enable-cloudbuild]
}

resource "google_project_iam_member" "default-cb-con-dev" {
  role    = "roles/container.developer"
  member = "serviceAccount:${data.google_project.project.number}@cloudbuild.gserviceaccount.com"
  depends_on = [google_project_service.project-enable-cloudbuild]
}

resource "google_project_iam_member" "default-cb-kms-cry-decry" {
  role    = "roles/cloudkms.cryptoKeyDecrypter"
  member = "serviceAccount:${data.google_project.project.number}@cloudbuild.gserviceaccount.com"
  depends_on = [google_project_service.project-enable-cloudbuild]
}

resource "google_project_iam_member" "default-cb-svcman-adm" {
  role    = "roles/servicemanagement.admin"
  member = "serviceAccount:${data.google_project.project.number}@cloudbuild.gserviceaccount.com"
  depends_on = [google_project_service.project-enable-cloudbuild]
}
