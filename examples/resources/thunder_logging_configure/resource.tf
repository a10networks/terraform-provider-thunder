provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_configure" "thunder_logging_configure" {
  line_feed_in_udp_syslog = 0
}