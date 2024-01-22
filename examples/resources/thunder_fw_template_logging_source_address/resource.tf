provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_template_logging_source_address" "thunder_fw_template_logging_source_address" {

  name = "test"
  ip   = "10.10.10.10"
  ipv6 = "2001:db8:0:200::1"
}