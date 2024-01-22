provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_server_radius" "thunder_aam_authentication_server_radius" {
  instance_list {
    name = "test"
    host {
      hostip = "10.10.10.10"
    }
    secret          = 1
    secret_string   = "108"
    port            = 1812
    interval        = 3
    retry           = 5
    accounting_port = 1813
    auth_type       = "pap"
    sampling_enable {
      counters1 = "all"
    }
  }
  sampling_enable {
    counters1 = "all"
  }
}