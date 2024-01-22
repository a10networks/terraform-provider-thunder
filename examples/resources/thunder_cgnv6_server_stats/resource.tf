provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_server_stats" "thunder_cgnv6_server_stats" {

  name = "test12345"
}
output "get_cgnv6_server_stats" {
  value = ["${data.thunder_cgnv6_server_stats.thunder_cgnv6_server_stats}"]
}