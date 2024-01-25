provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_template_client_ssl_stats" "thunder_slb_template_client_ssl_stats" {

  name = "temp1"
}
output "get_slb_template_client_ssl_stats" {
  value = ["${data.thunder_slb_template_client_ssl_stats.thunder_slb_template_client_ssl_stats}"]
}