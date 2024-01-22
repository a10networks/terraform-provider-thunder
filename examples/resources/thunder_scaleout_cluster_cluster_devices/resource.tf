provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_cluster_devices" "thunder_scaleout_cluster_cluster_devices" {

  cluster_discovery_timeout {
  }
  device_id_list {
    action = "enable"
  }
  enable = 0
  minimum_nodes {
    minimum_nodes_num = 11
  }
  cluster_id = 2
}