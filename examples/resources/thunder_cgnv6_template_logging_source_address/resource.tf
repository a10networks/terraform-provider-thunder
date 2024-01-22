provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_template_logging_source_address" "thunder_cgnv6_template_logging_source_address" {

  name = "test"
  ip   = "10.10.10.10"
}