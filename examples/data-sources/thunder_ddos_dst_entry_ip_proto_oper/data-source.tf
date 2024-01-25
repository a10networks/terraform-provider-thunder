provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_entry_ip_proto_oper" "thunder_ddos_dst_entry_ip_proto_oper" {
  dst_entry_name = "test"
  port_num       = 50
}
output "get_ddos_dst_entry_ip_proto_oper" {
  value = ["${data.thunder_ddos_dst_entry_ip_proto_oper.thunder_ddos_dst_entry_ip_proto_oper}"]
}