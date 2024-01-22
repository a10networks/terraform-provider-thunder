provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_vlan" "thunder_network_vlan" {
  sampling_enable {
    counters1 = "all"
  }
  shared_vlan = 0
  user_tag    = "38"
  vlan_num    = 2730
}