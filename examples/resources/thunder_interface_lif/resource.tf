provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_lif" "thunder_interface_lif" {
  ifname = 2
  ip {
    address_list {
      ipv4_address = "56.56.56.56"
      ipv4_netmask = "/16"
    }
  }
}