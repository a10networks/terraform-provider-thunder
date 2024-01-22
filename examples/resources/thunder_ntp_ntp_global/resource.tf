provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ntp_ntp_global" "thunder_ntp_ntp_global" {
  allow_data_ports = 1
}