provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_network_synchronization" "thunder_router_bgp_network_synchronization" {


  as_number               = "300"
  network_synchronization = 1
}