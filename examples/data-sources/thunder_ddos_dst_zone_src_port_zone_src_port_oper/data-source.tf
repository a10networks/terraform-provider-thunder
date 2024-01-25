provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_src_port_zone_src_port_oper" "thunder_ddos_dst_zone_src_port_zone_src_port_oper" {
  zone_name = "test"
  port_num  = 20
  protocol  = "tcp"

}
output "get_ddos_dst_zone_src_port_zone_src_port_oper" {
  value = ["${data.thunder_ddos_dst_zone_src_port_zone_src_port_oper.thunder_ddos_dst_zone_src_port_zone_src_port_oper}"]
}