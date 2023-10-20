provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_engine_id" "Snmp_Server_EngineID_Test" {
  eng_id = "ABABABABAB"
}