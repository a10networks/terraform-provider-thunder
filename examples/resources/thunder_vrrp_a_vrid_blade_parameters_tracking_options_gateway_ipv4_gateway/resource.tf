provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_vrid_blade_parameters_tracking_options_gateway_ipv4_gateway" "thunder_vrrp_a_vrid_blade_parameters_tracking_options_gateway_ipv4_gateway" {

  vrid_val      = 3
  ip_address    = "20.20.20.20"
  priority_cost = 22
}