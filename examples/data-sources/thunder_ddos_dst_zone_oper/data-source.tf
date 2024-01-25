provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_oper" "thunder_ddos_dst_zone_oper" {
  zone_name = "test"

}
output "get_ddos_dst_zone_oper" {
  value = ["${data.thunder_ddos_dst_zone_oper.thunder_ddos_dst_zone_oper}"]
}