provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scaleout_debug_ipv6_reachability_oper" "thunder_scaleout_debug_ipv6_reachability_oper" {

}
output "get_scaleout_debug_ipv6_reachability_oper" {
  value = ["${data.thunder_scaleout_debug_ipv6_reachability_oper.thunder_scaleout_debug_ipv6_reachability_oper}"]
}