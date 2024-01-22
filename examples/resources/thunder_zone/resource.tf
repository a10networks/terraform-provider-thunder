provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_zone" "thunder_zone" {

  name = "test_zone"
  interface {
    ethernet_list {
      interface_ethernet_start = 2
      interface_ethernet_end   = 2
    }
  }
  user_tag = "testing_zone"
  vlan {
    vlan_list {
      vlan_start = 2
      vlan_end   = 2
    }
  }
}