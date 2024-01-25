provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_src_entry_oper" "thunder_ddos_src_entry_oper" {
  src_entry_name = "2"

}
output "get_ddos_src_entry_oper" {
  value = ["${data.thunder_ddos_src_entry_oper.thunder_ddos_src_entry_oper}"]
}