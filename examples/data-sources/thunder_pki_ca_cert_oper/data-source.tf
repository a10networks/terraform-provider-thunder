provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_pki_ca_cert_oper" "thunder_pki_ca_cert_oper" {
}
output "get_pki_ca_cert_oper" {
  value = ["${data.thunder_pki_ca_cert_oper.thunder_pki_ca_cert_oper}"]
}