provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_nat_inside_source_list" "thunder_ipv6_nat_inside_source_list" {

  list_name = "a11"
  pool      = "K"
}