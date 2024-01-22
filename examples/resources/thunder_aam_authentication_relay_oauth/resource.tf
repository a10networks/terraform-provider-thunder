provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_relay_oauth" "thunder_aam_authentication_relay_oauth" {
  all        = 1
  name       = "3"
  relay_type = "access-token"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "48"
}