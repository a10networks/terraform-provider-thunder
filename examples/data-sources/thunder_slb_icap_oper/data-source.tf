provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_icap_oper" "thunder_slb_icap_oper" {
}
output "get_slb_icap_oper" {
  value = ["${data.thunder_slb_icap_oper.thunder_slb_icap_oper}"]
}