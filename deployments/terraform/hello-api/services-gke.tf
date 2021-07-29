
//resource "google_compute_address" "hello-api" {
resource "google_compute_global_address" "hello-api" {
  provider     = google-beta
  name         = "hello-api-address"
  project      = var.project_id
  //region       = var.region

  address_type = "EXTERNAL"
  labels = {
    app       = var.app_name
    component = "${var.app_name}-address"
  }
}

# load balancer
/*resource "kubernetes_service" "hello-api-lb" {
  metadata {
    name      = "${var.app_name}-lb"
    namespace = var.namespace
    labels = {
      app       = var.app_name
      component = "${var.app_name}-lb"
    }
  }
  spec {
    type = "LoadBalancer"
    selector = {
      app       = var.app_name
      component = "${var.app_name}-deploy"
    }
    load_balancer_ip = google_compute_address.hello-api.address
    port {
      port = 443
      target_port = 9000
    }
  }
}*/

# nodePort + ingress
resource "kubernetes_service" "hello-api-node-port" {
  metadata {
    name      = "${var.app_name}-nodeport"
    namespace = var.namespace
    labels = {
      app       = var.app_name
      component = "${var.app_name}-nodeport"
    }
  }
  spec {
    type = "NodePort"
    selector = {
      app       = var.app_name
      component = "${var.app_name}-deploy"
    }
    port {
      port = 9000
      protocol = "TCP"
      target_port = 9000
    }
  }
}
resource "kubernetes_ingress" "hello-api-ingress" {
  metadata {
    name = "${var.app_name}-ingress"
    namespace = var.namespace
    labels = {
      app       = var.app_name
      component = "${var.app_name}-ingress"
    }
    annotations = {
      // Not work with regional static ip
      "kubernetes.io/ingress.global-static-ip-name": google_compute_global_address.hello-api.name
      //"kubernetes.io/ingress.regional-static-ip-name": google_compute_address.hello-api.name
    }
  }
  spec {
    backend {
      service_name = kubernetes_service.hello-api-node-port.metadata[0].name
      service_port = kubernetes_service.hello-api-node-port.spec[0].port[0].port
    }
  }
}
