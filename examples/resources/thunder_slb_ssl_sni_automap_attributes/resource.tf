provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_ssl_sni_automap_attributes" "thunder_slb_ssl_sni_automap_attributes" {
  sni_delete_factor = 234
  sni_lower_limit   = 1231
  sni_upper_limit   = 3422
}