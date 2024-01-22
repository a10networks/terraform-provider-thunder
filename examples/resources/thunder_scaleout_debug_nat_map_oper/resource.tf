provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scaleout_debug_nat_map_oper" "thunder_scaleout_debug_nat_map_oper" {
}
output "get_scaleout_debug_nat_map_oper" {
  value = ["${data.thunder_scaleout_debug_nat_map_oper.thunder_scaleout_debug_nat_map_oper}"]
}