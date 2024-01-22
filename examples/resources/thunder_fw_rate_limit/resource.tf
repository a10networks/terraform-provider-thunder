provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_rate_limit" "thunder_fw_rate_limit" {
  summary {
  }
}