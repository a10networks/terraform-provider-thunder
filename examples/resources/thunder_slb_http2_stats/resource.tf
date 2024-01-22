provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_http2_stats" "thunder_slb_http2_stats" {
}
output "get_slb_http2_stats" {
  value = ["${data.thunder_slb_http2_stats.thunder_slb_http2_stats}"]
}