provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_lsn_port_unavailable" "thunder_logging_lsn_port_unavailable" {
  ip_based = 0
}