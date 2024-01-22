provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_port_list" "thunder_cgnv6_port_list" {
  name = "test"
  port_config {
    original_port   = 55532
    translated_port = 26432
  }
  user_tag = "test"
}