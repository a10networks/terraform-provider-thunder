provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_snmpv3_user" "Snmp_Server_SNMPv3_User_Test" {
  username          = "testing"
  group             = "abc"
  v3                = "auth"
  auth_val          = "md5"
  passwd            = "pass123"
  pw_encrypted      = "encpass123"
  priv              = "aes"
  encpasswd         = "pass123"
  priv_pw_encrypted = "encpass123"
}