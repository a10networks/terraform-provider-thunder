provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_cluster_devices_minimum_nodes" "thunder_scaleout_cluster_cluster_devices_minimum_nodes" {

  cluster_id        = 2
  minimum_nodes_num = 11
}