provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_route_rib" "rib_3" {
  ip_dest_addr = "7.7.0.0"
  ip_mask      = "/16"
  ip_nexthop_lif {
    lif                     = "Lif1"
    description_nexthop_lif = "Lif_Descrition"
  }
}