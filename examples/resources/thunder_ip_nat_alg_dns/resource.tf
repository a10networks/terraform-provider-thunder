provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_alg_dns" "thunder_ip_nat_alg_dns" {
  dns_alg = "disable"
}