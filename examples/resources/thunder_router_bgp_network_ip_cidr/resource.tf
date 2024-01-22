provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_network_ip_cidr" "thunder_router_bgp_network_ip_cidr" {

  as_number         = "300"
  network_ipv4_cidr = "19.0.2.0/24"
}