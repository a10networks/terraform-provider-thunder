provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_local_device" "thunder_scaleout_cluster_local_device" {

  cluster_id  = 2
  action      = "enable"
  priority    = 209
  start_delay = 167
}