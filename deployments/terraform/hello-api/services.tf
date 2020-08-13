
resource "kubernetes_service" "hello-api-lb" {
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

    port {
      port = 80
      target_port = 8080
    }
  }
}
