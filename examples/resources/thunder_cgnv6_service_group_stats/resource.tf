provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_service_group_stats" "thunder_cgnv6_service_group_stats" {

  name = "testing"
}
output "get_cgnv6_service_group_stats" {
  value = ["${data.thunder_cgnv6_service_group_stats.thunder_cgnv6_service_group_stats}"]
}