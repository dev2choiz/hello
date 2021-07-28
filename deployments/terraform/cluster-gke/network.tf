
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

## managed by google
/*resource "google_compute_managed_ssl_certificate" "managed-certificate" {
  description = "${var.domain} certificate"
  name = "dev2choiz-cert"
  managed {
    domains = [
      "${var.domain}.",
      "hello.${var.domain}.",
    ]
  }
}*/
