provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_relay_ntlm" "thunder_aam_authentication_relay_ntlm" {
  domain                = "14"
  large_request_disable = 0
  name                  = "test"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "107"
  version  = 2
}