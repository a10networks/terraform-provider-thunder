provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_local_device_l2_redirect" "thunder_scaleout_cluster_local_device_l2_redirect" {

  cluster_id   = 2
  redirect_eth = 2
}