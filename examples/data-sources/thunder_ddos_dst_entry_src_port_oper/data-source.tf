provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_entry_src_port_oper" "thunder_ddos_dst_entry_src_port_oper" {
  dst_entry_name = "test"
  port_num       = 8453
  protocol       = "dns-udp"

}
output "get_ddos_dst_entry_src_port_oper" {
  value = ["${data.thunder_ddos_dst_entry_src_port_oper.thunder_ddos_dst_entry_src_port_oper}"]
}