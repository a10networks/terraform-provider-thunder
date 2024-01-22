provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_relay_kerberos" "thunder_aam_authentication_relay_kerberos" {
  instance_list {
    name             = "test"
    kerberos_realm   = "17"
    kerberos_kdc     = "3"
    kerberos_account = "5"
    password         = 1
    secret_string    = "63"
    port             = 88
    timeout          = 10
    sampling_enable {
      counters1 = "all"
    }
  }
  sampling_enable {
    counters1 = "all"
  }
}