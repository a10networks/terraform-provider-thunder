provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_pki_cert_stats_oper" "thunder_pki_cert_stats_oper" {
}
output "get_pki_cert_stats_oper" {
  value = ["${data.thunder_pki_cert_stats_oper.thunder_pki_cert_stats_oper}"]
}