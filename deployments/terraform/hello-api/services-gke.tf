
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

output "hello-api-address" {
  value = google_compute_global_address.hello-api.address
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

# BackendConfig + FrontConfig + nodePort + ingress
resource "kubernetes_manifest" "hello-api-backend-config" {
  provider = kubernetes-alpha
  manifest = {
    apiVersion = "cloud.google.com/v1"
    kind = "BackendConfig"
    metadata = {
      name      = "${var.app_name}-backend-config"
      namespace = var.namespace
    }
    spec = {
      healthCheck = {
        checkIntervalSec = 5
        timeoutSec = 2
        healthyThreshold = 1
        unhealthyThreshold = 10
        type = "HTTP"
        requestPath = "/healthz"
        port = kubernetes_service.hello-api-node-port.spec.0.port.0.node_port
      }
    }
  }
}

resource "kubernetes_manifest" "hello-api-frontend-config" {
  provider = kubernetes-alpha
  manifest = {
    apiVersion = "networking.gke.io/v1beta1"
    kind = "FrontendConfig"
    metadata = {
      name      = "${var.app_name}-frontend-config"
      namespace = var.namespace
    }
    spec = {
      redirectToHttps = {
        enabled = true
      }
    }
  }
}

resource "kubernetes_service" "hello-api-node-port" {
  metadata {
    name      = "${var.app_name}-nodeport"
    namespace = var.namespace
    labels = {
      app       = var.app_name
      component = "${var.app_name}-nodeport"
    }
    annotations = {
      "cloud.google.com/backend-config": "{\"default\":\"${var.app_name}-backend-config\"}"
    }
  }
  spec {
    type = "NodePort"
    selector = {
      app       = var.app_name
      component = "${var.app_name}-deploy"
    }
    port {
      protocol = "TCP"
      port = 9001
      target_port = 9000 # esp
      //target_port = 8080 # api
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
      "networking.gke.io/v1beta1.FrontendConfig": "${var.app_name}-frontend-config"
    }
  }
  spec {
    tls {
      secret_name = kubernetes_secret.hello-tls-credentials.metadata[0].name
    }
    backend {
      service_name = kubernetes_service.hello-api-node-port.metadata[0].name
      service_port = kubernetes_service.hello-api-node-port.spec[0].port[0].port
    }
  }
}
