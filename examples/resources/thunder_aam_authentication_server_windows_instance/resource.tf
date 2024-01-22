provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_server_windows_instance" "thunder_aam_authentication_server_windows_instance" {
  auth_protocol {
    ntlm_disable                  = 0
    ntlm_version                  = 2
    kerberos_disable              = 1
    kerberos_port                 = 88
    kerberos_password_change_port = 464
    kdc_validate                  = 1
    kerberos_kdc_validation {
      kdc_spn      = "15"
      kdc_account  = "66"
      kdc_password = 1
      kdc_pwd      = "11"
    }
  }
  host {
    hostip = "10.10.10.10"
  }
  name  = "test"
  realm = "29"
  sampling_enable {
    counters1 = "all"
  }
  support_apacheds_kdc = 0
  timeout              = 10
}