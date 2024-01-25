provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_fixed_nat_port_mapping_files_oper" "thunder_cgnv6_fixed_nat_port_mapping_files_oper" {

}
output "get_cgnv6_fixed_nat_port_mapping_files_oper" {
  value = ["${data.thunder_cgnv6_fixed_nat_port_mapping_files_oper.thunder_cgnv6_fixed_nat_port_mapping_files_oper}"]
}