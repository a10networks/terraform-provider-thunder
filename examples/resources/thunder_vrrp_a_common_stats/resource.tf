provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vrrp_a_common_stats" "thunder_vrrp_a_common_stats" {

}
output "get_vrrp_a_common_stats" {
  value = ["${data.thunder_vrrp_a_common_stats.thunder_vrrp_a_common_stats}"]
}