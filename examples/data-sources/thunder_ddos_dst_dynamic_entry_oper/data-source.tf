provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_dynamic_entry_oper" "thunder_ddos_dst_dynamic_entry_oper" {
}
output "get_ddos_dst_dynamic_entry_oper" {
  value = ["${data.thunder_ddos_dst_dynamic_entry_oper.thunder_ddos_dst_dynamic_entry_oper}"]
}