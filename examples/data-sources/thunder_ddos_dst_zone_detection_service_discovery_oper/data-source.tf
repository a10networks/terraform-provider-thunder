provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_detection_service_discovery_oper" "thunder_ddos_dst_zone_detection_service_discovery_oper" {
  zone_name = "test"

}
output "get_ddos_dst_zone_detection_service_discovery_oper" {
  value = ["${data.thunder_ddos_dst_zone_detection_service_discovery_oper.thunder_ddos_dst_zone_detection_service_discovery_oper}"]
}