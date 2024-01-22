provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_cloud_integration_ecosystem_consul" "thunder_acos_cloud_integration_ecosystem_consul" {
  action                = "disable"
  health_check_interval = "5"
  host_name             = "67"
  port                  = 277
}