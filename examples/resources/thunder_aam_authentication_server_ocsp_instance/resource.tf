provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_server_ocsp_instance" "thunder_aam_authentication_server_ocsp_instance" {
  http_version = 1
  name         = "61"
  sampling_enable {
    counters1 = "all"
  }
  url          = "http://192.168.0.1:80/"
  version_type = "1.1"
}