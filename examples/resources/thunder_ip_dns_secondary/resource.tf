provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_dns_secondary" "thunder_ip_dns_secondary" {

  ip_v4_addr = "9.9.9.9"
}