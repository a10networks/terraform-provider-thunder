provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_local_device_traffic_redirection_reachability_options" "thunder_scaleout_cluster_local_device_traffic_redirection_reachability_options" {

  cluster_id         = 2
  skip_default_route = 1
}