provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_lif_ipv6_router_ospf" "thunder_interface_lif_ipv6_router_ospf" {

  ifname = "test"
  area_list {
    area_id_num  = 0
    area_id_addr = "10.10.10.19"
    tag          = "46"
    instance_id  = 0
  }
}