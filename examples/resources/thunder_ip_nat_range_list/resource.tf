provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_range_list" "thunder_ip_nat_range_list" {
  name                   = 11
  global_start_ipv4_addr = "10.0.2.6"
  global_netmaskv4       = "255.255.255.255"
  local_start_ipv4_addr  = "10.0.2.7"
  local_netmaskv4        = "255.255.255.255"
  v4_count               = 1
}