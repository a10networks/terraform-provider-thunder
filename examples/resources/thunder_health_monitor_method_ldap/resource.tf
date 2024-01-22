provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_ldap" "thunder_health_monitor_method_ldap" {

  name                 = "tf_test"
  acceptnotfound       = 1
  acceptresref         = 1
  basedn               = "14"
  ldap                 = 1
  ldap_binddn          = "1lac"
  ldap_password        = 1
  ldap_password_string = "a30"
  ldap_port            = 38923
  ldap_query           = "52"
  ldap_run_search      = 1
  ldap_security        = "overssl"
}