provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_flowspec_icmp_code" "thunder_flowspec_icmp_code" {

  name                = "test"
  code                = 27
  icmp_code_attribute = "eq"
}