# out of scope but if need to provision server(aws) then terraform best suits

provider "kubernetes" {
  config_path = "~/.kube/config"
}

resource "kubernetes_namespace" "frontend_namespace" {
  metadata {
    name = "frontend"
  }
}

resource "kubernetes_namespace" "backend_namespace" {
  metadata {
    name = "backend"
  }
}

resource "kubernetes_deployment" "frontend" {
  metadata {
    name      = "frontend-deployment"
    namespace = kubernetes_namespace.frontend_namespace.metadata[0].name
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "frontend"
      }
    }

    template {
      metadata {
        labels = {
          app = "frontend"
        }
      }

      spec {
        container {
          name  = "frontend"
          image = "php:7.4-apache"
          ports {
            container_port = 80
          }
          volume_mount {
            name       = "frontend-files"
            mount_path = "/var/www/html/frontend"
          }
        }

        volume {
          name = "frontend-files"
          host_path {
            path = "C:\laragon\www\akc_test\frontend" 
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "frontend" {
  metadata {
    name      = "frontend-service"
    namespace = kubernetes_namespace.frontend_namespace.metadata[0].name
  }

  spec {
    selector = {
      app = "frontend"
    }

    port {
      port        = 80
      target_port = 80
    }

    type = "LoadBalancer"
  }
}

resource "kubernetes_deployment" "backend" {
  metadata {
    name      = "backend-deployment"
    namespace = kubernetes_namespace.backend_namespace.metadata[0].name
  }

  # Define the backend deployment similar to the frontend deployment...
  # (Replicas, container image, ports, etc.)
  # Use the appropriate values for the backend container image and ports.
}

resource "kubernetes_service" "backend" {
  metadata {
    name      = "backend-service"
    namespace = kubernetes_namespace.backend_namespace.metadata[0].name
  }

  # Define the backend service similar to the frontend service...
  # (Port mappings, type, etc.)
}
