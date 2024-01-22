provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_default_app_type" "thunder_ddos_src_default_app_type" {
  default_address_type = "ip"
  protocol             = "dns"
  user_tag             = "32"
}