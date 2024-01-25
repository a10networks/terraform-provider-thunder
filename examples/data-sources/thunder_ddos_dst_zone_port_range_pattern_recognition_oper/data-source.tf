provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_port_range_pattern_recognition_oper" "thunder_ddos_dst_zone_port_range_pattern_recognition_oper" {
  zone_name        = "test"
  port_range_end   = 22
  port_range_start = 20
  protocol         = "tcp"

}
output "get_ddos_dst_zone_port_range_pattern_recognition_oper" {
  value = ["${data.thunder_ddos_dst_zone_port_range_pattern_recognition_oper.thunder_ddos_dst_zone_port_range_pattern_recognition_oper}"]
}