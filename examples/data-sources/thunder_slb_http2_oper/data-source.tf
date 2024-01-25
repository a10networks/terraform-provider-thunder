provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_http2_oper" "thunder_slb_http2_oper" {
}
output "get_slb_http2_oper" {
  value = ["${data.thunder_slb_http2_oper.thunder_slb_http2_oper}"]
}