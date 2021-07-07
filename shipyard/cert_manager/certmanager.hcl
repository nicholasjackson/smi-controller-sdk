k8s_config "cert-manager-local" {
  cluster = "k8s_cluster.dc1"
  
  paths = [
    "${file_dir()}/cert-manager-crds.yaml",
  ]

  wait_until_ready = true
}

helm "cert-manager-local" {
  depends_on = ["k8s_config.cert-manager-local"]

  create_namespace = true
  namespace = "smi"
  cluster = "k8s_cluster.dc1"

  chart = "github.com/jetstack/cert-manager?ref=v1.2.0/deploy/charts//cert-manager"

  values = "${file_dir()}/helm-values.yaml" 
}