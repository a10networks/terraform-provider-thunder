provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_loopback_ipv6_router_ospf" "thunder_interface_loopback_ipv6_router_ospf" {

  ifnum = 4
  area_list {
    area_id_num  = 0
    area_id_addr = "10.10.10.10"
    tag          = "96"
    instance_id  = 0
  }
}