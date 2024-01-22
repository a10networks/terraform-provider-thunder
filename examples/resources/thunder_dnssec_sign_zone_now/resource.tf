provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_dnssec_sign_zone_now" "thunder_dnssec_sign_zone_now" {
  zone_name = ""
}