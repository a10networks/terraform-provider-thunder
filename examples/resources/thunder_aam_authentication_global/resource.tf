provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_global" "thunder_aam_authentication_global" {
  max_auth_resp_size = 65536
  sampling_enable {
    counters1 = "all"
  }
}