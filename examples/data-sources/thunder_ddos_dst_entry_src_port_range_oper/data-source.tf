provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_entry_src_port_range_oper" "thunder_ddos_dst_entry_src_port_range_oper" {
  dst_entry_name       = "test"
  protocol             = "udp"
  src_port_range_end   = 24
  src_port_range_start = 20

}
output "get_ddos_dst_entry_src_port_range_oper" {
  value = ["${data.thunder_ddos_dst_entry_src_port_range_oper.thunder_ddos_dst_entry_src_port_range_oper}"]
}