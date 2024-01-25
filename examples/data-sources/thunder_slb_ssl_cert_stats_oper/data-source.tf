provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_ssl_cert_stats_oper" "thunder_slb_ssl_cert_stats_oper" {
}
output "get_slb_ssl_cert_stats_oper" {
  value = ["${data.thunder_slb_ssl_cert_stats_oper.thunder_slb_ssl_cert_stats_oper}"]
}