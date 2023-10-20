provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_active_rule_set" "test_thunder_fw_active_rule_set" {
  name               = "testing"
  session_aging      = "temp"
  override_nat_aging = 1
}