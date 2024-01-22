provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat64_prefix" "thunder_cgnv6_nat64_prefix" {

  prefix_val = "2001:db8:3333:4444:5555:6666:7777:8888/96"
  vrid       = 3

}