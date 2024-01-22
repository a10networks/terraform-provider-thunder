provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_zone_vlan" "thunder_zone_vlan" {

  name = "test_zone"
  vlan_list {
    vlan_start = 2
    vlan_end   = 2
  }
}