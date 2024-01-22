provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_zone_interface" "thunder_zone_interface" {

  name = "test_zone"
  ethernet_list {
    interface_ethernet_start = 2
    interface_ethernet_end   = 2
  }
}