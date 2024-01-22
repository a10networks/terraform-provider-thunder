provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_vrid_blade_parameters_tracking_options" "thunder_vrrp_a_vrid_blade_parameters_tracking_options" {

  vrid_val = 3
  interface {
    ethernet      = 2
    priority_cost = 22
  }
  bgp {
    bgp_ipv4_address_cfg {
      bgp_ipv4_address = "10.10.10.10"
      priority_cost    = 22
    }
  }
  gateway {
    ipv4_gateway_list {
      ip_address    = "20.20.20.20"
      priority_cost = 22
    }
  }
}