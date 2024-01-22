provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_active_rule_set" "thunder_fw_active_rule_set" {

  name               = "test"
  override_nat_aging = 1
}