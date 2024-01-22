provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_flowspec_icmp_type" "thunder_flowspec_icmp_type" {

  name                = "test"
  icmp_type_attribute = "eq"
  type                = 19
}