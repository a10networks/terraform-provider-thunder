provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_vlan_global" "thunder_network_vlan_global" {
  enable_def_vlan_l2_forwarding = 0
  l3_vlan_fwd_disable           = 0
  sampling_enable {
    counters1 = "all"
  }
}