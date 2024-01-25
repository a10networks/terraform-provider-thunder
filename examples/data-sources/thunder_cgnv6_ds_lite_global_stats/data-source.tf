provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_ds_lite_global_stats" "thunder_cgnv6_ds_lite_global_stats" {
}
output "get_cgnv6_ds_lite_global_stats" {
  value = ["${data.thunder_cgnv6_ds_lite_global_stats.thunder_cgnv6_ds_lite_global_stats}"]
}