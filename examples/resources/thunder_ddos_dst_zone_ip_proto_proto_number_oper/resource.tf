provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_ip_proto_proto_number_oper" "thunder_ddos_dst_zone_ip_proto_proto_number_oper" {
  zone_name    = "test"
  protocol_num = "80"

}
output "get_ddos_dst_zone_ip_proto_proto_number_oper" {
  value = ["${data.thunder_ddos_dst_zone_ip_proto_proto_number_oper.thunder_ddos_dst_zone_ip_proto_proto_number_oper}"]
}