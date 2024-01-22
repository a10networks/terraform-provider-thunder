provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_scaleout_address_mapping_oper" "thunder_cgnv6_scaleout_address_mapping_oper" {
}
output "get_cgnv6_scaleout_address_mapping_oper" {
  value = ["${data.thunder_cgnv6_scaleout_address_mapping_oper.thunder_cgnv6_scaleout_address_mapping_oper}"]
}