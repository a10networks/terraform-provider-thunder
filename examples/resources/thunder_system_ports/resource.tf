provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_ports" "System_Ports" {
  link_detection_interval = 50
}