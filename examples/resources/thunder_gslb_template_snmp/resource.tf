provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_gslb_template_snmp" "thunder_gslb_template_snmp" {
  snmp_name = "testSNMP"
  version = "v3"
  community = "test"
  security_level = "no-auth"
  oid = "testOID"
  interface = 2
  username = "test_user"
  auth_key = "test"
  priv_key = "testing"
  host = "test.com"
  port = 233
  interval = 22
  auth_proto = "sha"
  priv_proto = "des"
  context_name = "testContext"
  context_engine_id = "2"
  security_engine_id = "4"
  user_tag = "suraj"
}