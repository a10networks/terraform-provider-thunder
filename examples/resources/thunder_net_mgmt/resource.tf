provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_net_mgmt" "thunder_net_mgmt" {
  snmp {
    engineid {
    }
    stats {
    }
  }
}