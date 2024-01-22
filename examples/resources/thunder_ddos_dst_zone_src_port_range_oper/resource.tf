provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_src_port_range_oper" "thunder_ddos_dst_zone_src_port_range_oper" {
  zone_name            = "test"
  src_port_range_start = 20
  src_port_range_end   = 22
  protocol             = "tcp"

}
output "get_ddos_dst_zone_src_port_range_oper" {
  value = ["${data.thunder_ddos_dst_zone_src_port_range_oper.thunder_ddos_dst_zone_src_port_range_oper}"]
}