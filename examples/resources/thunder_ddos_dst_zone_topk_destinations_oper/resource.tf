provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_topk_destinations_oper" "thunder_ddos_dst_zone_topk_destinations_oper" {

  zone_name = "testZone"

}
output "get_ddos_dst_zone_topk_destinations_oper" {
  value = ["${data.thunder_ddos_dst_zone_topk_destinations_oper.thunder_ddos_dst_zone_topk_destinations_oper}"]
}