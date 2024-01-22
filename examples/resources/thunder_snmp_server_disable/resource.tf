provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_disable" "thunder_snmp_server_disable" {
  a10cmsubagent = 0
  traps {
    all        = 0
    vrrp_a     = 0
    slb        = 0
    slb_change = 0
  }
}