provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_other_zone_ipproto_stats" "thunder_ddos_other_zone_ipproto_stats" {

}
output "get_ddos_other_zone_ipproto_stats" {
  value = ["${data.thunder_ddos_other_zone_ipproto_stats.thunder_ddos_other_zone_ipproto_stats}"]
}