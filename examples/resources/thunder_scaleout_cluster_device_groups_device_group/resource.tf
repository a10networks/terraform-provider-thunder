provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_device_groups_device_group" "thunder_scaleout_cluster_device_groups_device_group" {

  cluster_id   = 2
  device_group = 2
  device_id_list {
    device_id_start = 11
    device_id_end   = 12
  }
  user_tag = "8"
}