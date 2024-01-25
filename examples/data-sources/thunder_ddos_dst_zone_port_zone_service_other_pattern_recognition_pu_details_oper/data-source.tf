provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_port_zone_service_other_pattern_recognition_pu_details_oper" "thunder_ddos_dst_zone_port_zone_service_other_pattern_recognition_pu_details_oper" {
  zone_name  = "test"
  port_other = "other"
  protocol   = "tcp"

}
output "get_ddos_dst_zone_port_zone_service_other_pattern_recognition_pu_details_oper" {
  value = ["${data.thunder_ddos_dst_zone_port_zone_service_other_pattern_recognition_pu_details_oper.thunder_ddos_dst_zone_port_zone_service_other_pattern_recognition_pu_details_oper}"]
}