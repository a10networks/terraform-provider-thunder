provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_relay_http_basic_instance" "thunder_aam_authentication_relay_http_basic_instance" {
  domain        = "test"
  domain_format = "down-level-logon-name"
  name          = "test"
  sampling_enable {
    counters1 = "all"
  }
}