
data "google_dns_managed_zone" "main-zone" {
  project = var.project_id
  name    = "main-zone"
}

resource "google_dns_record_set" "resource-recordset" {
  provider     = google-beta
  project      = var.project_id
  managed_zone = data.google_dns_managed_zone.main-zone.name
  name         = data.google_dns_managed_zone.main-zone.dns_name
  type         = "A"
  rrdatas      = [google_compute_global_address.hello-api.address]
  #rrdatas      = [google_compute_address.hello-api.address]
  ttl          = 300
}

resource "google_dns_record_set" "resource-recordset-www" {
  provider     = google-beta
  project      = var.project_id
  managed_zone = data.google_dns_managed_zone.main-zone.name
  name         = "www.${data.google_dns_managed_zone.main-zone.dns_name}"
  type         = "CNAME"
  rrdatas      = ["${var.domain}."]
  ttl          = 300
}

resource "google_dns_record_set" "resource-recordset-sub" {
  provider     = google-beta
  project      = var.project_id
  managed_zone = data.google_dns_managed_zone.main-zone.name
  name         = "hello.${data.google_dns_managed_zone.main-zone.dns_name}"
  type         = "CNAME"
  rrdatas      = ["${var.domain}."]
  ttl          = 300
}

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
  data = {
    "tls.crt" = file(var.certCrtPath)
    "tls.key" = file(var.certPrivateKeyPath)
  }
}
