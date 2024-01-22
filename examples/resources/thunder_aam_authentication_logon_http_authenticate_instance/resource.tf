provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_logon_http_authenticate_instance" "thunder_aam_authentication_logon_http_authenticate_instance" {
  account_lock = 0
  auth_method {
    basic {
      basic_realm = "17"
    }
    ntlm {
      ntlm_enable = 1
    }
    negotiate {
      negotiate_enable = 0
    }
  }
  duration = 1800
  name     = "tetst"
  retry    = 3
  sampling_enable {
    counters1 = "all"
  }
}