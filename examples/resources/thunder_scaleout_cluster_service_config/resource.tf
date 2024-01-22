provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_service_config" "thunder_scaleout_cluster_service_config" {

  cluster_id = 2
  enable     = 0
  template_list {
    name         = "12"
    device_group = 3
    user_tag     = "19"
  }
}