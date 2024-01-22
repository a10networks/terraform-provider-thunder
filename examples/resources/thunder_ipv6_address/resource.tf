provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_address" "thunder_ipv6_address" {
  ipv6_address = "e1f1:bdd5:e42a:57df:cc7c:6dd5:fb8e:62c8/24"
  anycast      = 1
}