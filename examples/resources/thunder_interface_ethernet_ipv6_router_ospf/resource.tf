provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_ipv6_router_ospf" "thunder_interface_ethernet_ipv6_router_ospf" {
  ifnum = 2
  area_list {
    area_id_num  = 0
    area_id_addr = "10.11.2.5"
    tag          = "87"
    instance_id  = 0
  }
}