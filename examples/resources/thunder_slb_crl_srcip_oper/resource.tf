provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_crl_srcip_oper" "thunder_slb_crl_srcip_oper" {
}
output "get_slb_crl_srcip_oper" {
  value = ["${data.thunder_slb_crl_srcip_oper.thunder_slb_crl_srcip_oper}"]
}