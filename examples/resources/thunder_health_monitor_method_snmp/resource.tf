provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_health_monitor_method_snmp" "thunder_health_monitor_method_snmp" {

  name      = "tf_test"
  community = "public"
  oid {
    mib = "sysDescr"
  }
  operation {
    oper_type = "getnext"
  }
  snmp      = 1
  snmp_port = 16111
}