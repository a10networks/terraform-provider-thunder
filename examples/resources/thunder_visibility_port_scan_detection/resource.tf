provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_port_scan_detection" "thunder_visibility_port_scan_detection" {
  events   = 11
  interval = 66
}