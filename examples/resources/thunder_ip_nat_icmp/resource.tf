provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_icmp" "NatIcmp" {
  respond_to_ping          = 1
  always_source_nat_errors = 1
}