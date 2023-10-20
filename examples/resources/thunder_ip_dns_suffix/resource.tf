provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ip_dns_suffix" "testname" {
  domain_name = "A10test"
}