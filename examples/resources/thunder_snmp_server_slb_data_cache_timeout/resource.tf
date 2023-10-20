provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_slb_data_cache_timeout" "resourceSnmpServerSlbDataCacheTimeoutTest" {
  slblimit = 75
}