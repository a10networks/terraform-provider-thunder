provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_entry_oper" "thunder_ddos_dst_entry_oper" {
  dst_entry_name = "test"

}
output "get_ddos_dst_entry_oper" {
  value = ["${data.thunder_ddos_dst_entry_oper.thunder_ddos_dst_entry_oper}"]
}