provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_src_entry_app_type" "thunder_ddos_src_entry_app_type" {
  src_entry_name = "25"
  protocol       = "dns"
  user_tag       = "108"

}