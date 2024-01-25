provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_dst_zone_stats" "thunder_ddos_dst_zone_stats" {
  zone_name = "test"

}
output "get_ddos_dst_zone_stats" {
  value = ["${data.thunder_ddos_dst_zone_stats.thunder_ddos_dst_zone_stats}"]
}