provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_user" "resourceSnmpServerUserTest" {
  username  = "testing"
  auth_val  = "sha"
  v3        = "auth"
  group     = "abc"
  passwd    = "abcdefgh"
  priv      = "aes"
  encpasswd = "abcdefgh"
}
