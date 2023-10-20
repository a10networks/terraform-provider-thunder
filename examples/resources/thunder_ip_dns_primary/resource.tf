provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ip_dns_primary" "dnsPrimary" {
  ip_v4_addr = "10.10.10.10"
}