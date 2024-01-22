provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_clear_session_filter" "test_thunder_fw_clear_session_filter" {
  status = "enable"
}