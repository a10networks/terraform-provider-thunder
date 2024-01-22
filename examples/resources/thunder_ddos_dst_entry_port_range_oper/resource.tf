provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_entry_port_range_oper" "thunder_ddos_dst_entry_port_range_oper" {

  port_range_end   = 90
  port_range_start = 80
  protocol         = "tcp"
  dst_entry_name   = "test"
}
output "get_ddos_dst_entry_port_range_oper" {
  value = ["${data.thunder_ddos_dst_entry_port_range_oper.thunder_ddos_dst_entry_port_range_oper}"]
}