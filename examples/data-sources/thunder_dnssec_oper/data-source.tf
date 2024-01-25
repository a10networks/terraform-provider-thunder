provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_dnssec_oper" "thunder_dnssec_oper" {
}
output "get_dnssec_oper" {
  value = ["${data.thunder_dnssec_oper.thunder_dnssec_oper}"]
}