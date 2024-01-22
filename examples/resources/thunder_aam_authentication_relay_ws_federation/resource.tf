provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_relay_ws_federation" "thunder_aam_authentication_relay_ws_federation" {
  application_server = "sharepoint"
  authentication_uri = "17"
  name               = "test"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "thunder_aam_authentication_relay_ws_federation"
}