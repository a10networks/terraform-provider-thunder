provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_service_config_template" "thunder_scaleout_cluster_service_config_template" {

  cluster_id   = 2
  name         = "12"
  device_group = 3
  user_tag     = "19"
}