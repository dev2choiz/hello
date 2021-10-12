
resource "google_compute_global_address" "front_next_cr_address" {
    provider = google-beta
    name = "hello-front-next-cr-address"
    project = var.project_id

    address_type = "EXTERNAL"
    labels = {
        app = "${var.app_name}-front-next-cr"
        component = "${var.app_name}-front-next-cr-address"
    }
}

output "hello_front_next_cr_address" {
    value = google_compute_global_address.front_next_cr_address.address
}

resource "google_compute_backend_service" "front_next_cr_backend_service" {
    provider = google-beta
    name = "front-next-cr-backend-service"
    load_balancing_scheme = "EXTERNAL"
    protocol = "HTTPS"

    enable_cdn = true
    cdn_policy {
        default_ttl = 3600
        cache_mode = "CACHE_ALL_STATIC"
        signed_url_cache_max_age_sec = 3600
        cache_key_policy {
            include_query_string = true
        }
    }
    custom_response_headers = [
        "X-Cache-Status: {cdn_cache_status}",
        "X-Cache-ID: {cdn_cache_id}",
    ]

    backend {
        // add the NEG to the backend service
        group = google_compute_region_network_endpoint_group.front_next_cr_neg.id
    }
}

resource "google_compute_url_map" "front_next_cr_url_map" {
    provider = google-beta
    name = "front-next-cr-lb"
    default_service = google_compute_backend_service.front_next_cr_backend_service.id
}

resource "google_compute_target_https_proxy" "front_next_cr_target_https_proxy" {
    provider = google-beta
    name = "front-next-cr-target-https-proxy"
    url_map = google_compute_url_map.front_next_cr_url_map.id
    ssl_certificates = [google_compute_ssl_certificate.default.id]
}

resource "google_compute_global_forwarding_rule" "front_next_cr_forwarding_rule" {
    provider = google-beta
    name = "front-next-cr-forwarding-rule"
    load_balancing_scheme = "EXTERNAL"
    target = google_compute_target_https_proxy.front_next_cr_target_https_proxy.id
    ip_address = google_compute_global_address.front_next_cr_address.address
    port_range = "443"
}

// http redirect rule
resource "google_compute_url_map" "front_next_cr_url_map_http_redirect" {
    provider = google-beta
    name = "front-next-cr-lb-http-redirect"
    default_url_redirect {
        strip_query = false
        https_redirect = true
        redirect_response_code = "MOVED_PERMANENTLY_DEFAULT"
    }
}

resource "google_compute_target_http_proxy" "front_next_cr_target_http_proxy" {
    provider = google-beta
    name = "front-next-cr-target-http-proxy"
    url_map = google_compute_url_map.front_next_cr_url_map_http_redirect.id
}

resource "google_compute_global_forwarding_rule" "front_next_cr_forwarding_rule_http" {
    provider = google-beta
    name = "front-next-cr-forwarding-rule-http"
    load_balancing_scheme = "EXTERNAL"
    target = google_compute_target_http_proxy.front_next_cr_target_http_proxy.id
    ip_address = google_compute_global_address.front_next_cr_address.address
    port_range = "80"
}

// Get cloud run data (deployed in the ci/cd pipeline)
data "google_cloud_run_service" "front_next_cloud_run" {
    provider = google-beta
    name = "hello-front-next"
    location = var.region
}

// NEG for the cloud run service
resource "google_compute_region_network_endpoint_group" "front_next_cr_neg" {
    provider = google-beta
    name = "front-next-cr-neg"
    region = var.region

    network_endpoint_type = "SERVERLESS"
    cloud_run {
        service = data.google_cloud_run_service.front_next_cloud_run.name
    }
}
