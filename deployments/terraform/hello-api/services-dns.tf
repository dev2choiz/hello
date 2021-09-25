
data "google_dns_managed_zone" "main-zone" {
  project = var.project_id
  name    = "main-zone"
}

resource "google_dns_record_set" "resource-recordset-api" {
  provider     = google-beta
  project      = var.project_id
  managed_zone = data.google_dns_managed_zone.main-zone.name
  name         = "api.${data.google_dns_managed_zone.main-zone.dns_name}"
  type         = "A"
  rrdatas      = [google_compute_global_address.hello-api.address] // ip address used for the ingress api
  ttl          = 300
}

resource "google_dns_record_set" "resource-recordset-front-cr" {
  provider     = google-beta
  project      = var.project_id
  managed_zone = data.google_dns_managed_zone.main-zone.name
  name         = "front.${data.google_dns_managed_zone.main-zone.dns_name}"
  type         = "A"
  rrdatas      = [google_compute_global_address.front_cr_address.address]
  ttl          = 300
}

/*
// Disabled, waiting the certificate updating
resource "google_dns_record_set" "resource-recordset-front-cr-www" {
  provider     = google-beta
  project      = var.project_id
  managed_zone = data.google_dns_managed_zone.main-zone.name
  name         = "www.${google_dns_record_set.resource-recordset-front-cr.name}"
  type         = "CNAME"
  rrdatas      = [google_dns_record_set.resource-recordset-front-cr.name]
  ttl          = 300
}*/

# ssl self managed
resource "google_compute_ssl_certificate" "default" {
  name        = "ssl-cert-default"
  description = "${var.domain} self managed certificate"
  private_key = file(var.certPrivateKeyPath)
  certificate = file(var.certCrtPath)
}

resource "kubernetes_secret" "hello-tls-credentials" {
  metadata {
    name      = "hello-tls-credentials"
    namespace = var.namespace
  }
  type = "kubernetes.io/tls"
  data = {
    "tls.crt" = file(var.certCrtPath)
    "tls.key" = file(var.certPrivateKeyPath)
    "server.crt" = file(var.certCrtPath)
    "server.key" = file(var.certPrivateKeyPath)
  }
}
