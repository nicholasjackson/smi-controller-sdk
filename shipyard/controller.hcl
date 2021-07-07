variable "smi_controller_k8s_cluster" {
  default = "dc1"
}

variable "smi_controller_k8s_network" {
  default = "dc1"
}

variable "smi_controller_enabled" {
  default = false
}

variable "smi_controller_webhook_enabled" {
  default = false
}

variable "smi_controller_webhook_port" {
  default = 9443
}

variable "smi_controller_namespace" {
  default = "shipyard"
}

variable "smi_controller_additional_dns" {
  default = "smi-webhook.shipyard.svc"
}

variable "smi_controller_helm_chart" {
  default = "${file_dir()}/../helm/smi-controller"
}

variable "install_helm_chart" {
  default = true
}

module "smi-controller" {
  disabled = (var.install_helm_chart == false)

  #source = "/home/nicj/go/src/github.com/shipyard-run/blueprints/modules/kubernetes-smi-controller"
  source = "github.com/shipyard-run/blueprints/modules/kubernetes-smi-controller"
}

# Create an ingress which exposes the locally running webhook from kubernetes
ingress "smi-webhook" {
  source {
    driver = "k8s"
    
    config {
      cluster = "k8s_cluster.dc1"
      port = 9443
    }
  }
  
  destination {
    driver = "local"
    
    config { 
      address = "localhost"
      port = 9443
    }
  
  }
}
