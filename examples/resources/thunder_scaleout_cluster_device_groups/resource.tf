provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_device_groups" "thunder_scaleout_cluster_device_groups" {

  cluster_id = 2
  device_group_list {
    device_group = 6
    device_id_list {
      device_id_start = 11
      device_id_end   = 12
    }
    user_tag = "94"
  }
  enable = 0
}