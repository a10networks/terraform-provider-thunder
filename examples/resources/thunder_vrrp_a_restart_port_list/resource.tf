provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_restart_port_list" "thunder_vrrp_a_restart_port_list" {
  ethernet_cfg {
    flap_ethernet_start = 1
    flap_ethernet_end   = 2
  }
  vrid_list {
    vrid_val = 7
    ethernet_cfg {
      flap_ethernet_start = 1
      flap_ethernet_end   = 2
    }
    user_tag = "test"
  }
}