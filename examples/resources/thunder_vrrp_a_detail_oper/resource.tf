provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_vrrp_a_detail_oper" "thunder_vrrp_a_detail_oper" {
}
output "get_vrrp_a_detail_oper" {
  value = ["${data.thunder_vrrp_a_detail_oper.thunder_vrrp_a_detail_oper}"]
}