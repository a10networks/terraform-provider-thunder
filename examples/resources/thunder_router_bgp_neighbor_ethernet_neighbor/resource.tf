provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_router_bgp_neighbor_ethernet_neighbor" "ethernetBGP" {
  as_number       = "300"
  process_id      = 200
  ethernet        = 3
  unnumbered      = 1
  peer_group_name = "A10"
}