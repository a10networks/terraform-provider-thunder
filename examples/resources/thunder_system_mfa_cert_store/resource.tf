provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_mfa_cert_store" "thunder_system_mfa_cert_store" {
  cert_store_path = "56"
  passwd_string   = "72"
  protocol        = "tftp"
  username        = "14"
}