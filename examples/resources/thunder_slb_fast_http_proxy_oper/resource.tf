provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_fast_http_proxy_oper" "thunder_slb_fast_http_proxy_oper" {
}
output "get_slb_fast_http_proxy_oper" {
  value = ["${data.thunder_slb_fast_http_proxy_oper.thunder_slb_fast_http_proxy_oper}"]
}