provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_resource_tracking_cpu" "thunder_ddos_resource_tracking_cpu" {
  enable = 1
}