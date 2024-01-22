provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_ssl_cert_revoke" "thunder_slb_ssl_cert_revoke" {
  sampling_enable {
    counters1 = "all"
  }
}