provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_dnssec_dnskey" "thunder_dnssec_dnskey" {
  key_delete = 1
  zone_name  = ""
}