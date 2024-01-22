provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_dnssec_dnskey" "thunder_import_periodic_dnssec_dnskey" {
  dnssec_dnskey = "test"
  period        = 18461731
  use_mgmt_port = 1
  remote_file   = "test"
}