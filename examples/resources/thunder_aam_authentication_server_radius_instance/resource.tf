provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_server_radius_instance" "thunder_aam_authentication_server_radius_instance" {
  accounting_port = 1813
  auth_type       = "pap"
  host {
    hostip = "10.10.10.10"
  }
  interval = 3
  name     = "test"
  port     = 1812
  retry    = 5
  sampling_enable {
    counters1 = "all"
  }
  secret        = 1
  secret_string = "97"
}