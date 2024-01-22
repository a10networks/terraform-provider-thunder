provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_kerberos_kdc" "thunder_health_monitor_method_kerberos_kdc" {

  name = "tf_test"
  kerberos_cfg {
    kinit               = 1
    kinit_pricipal_name = "a71"
    kinit_password      = "3"
    kinit_kdc           = "18.0.2.1"
    tcp_only            = 1
  }
}