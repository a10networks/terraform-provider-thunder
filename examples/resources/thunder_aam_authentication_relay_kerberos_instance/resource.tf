provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_aam_authentication_relay_kerberos_instance" "thunder_aam_authentication_relay_kerberos_instance" {

  kerberos_account = "111"
  kerberos_kdc     = "10"
  kerberos_realm   = "30"
  name             = "33"
  password         = 1
  port             = 88
  sampling_enable {
    counters1 = "all"
  }
  secret_string = "53"
  timeout       = 10

}