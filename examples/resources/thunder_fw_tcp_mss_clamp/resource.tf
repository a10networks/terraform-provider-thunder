provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_tcp_mss_clamp" "test_thunder_fw_tcp_mss_clamp" {
  mss_clamp_type = "subtract"
  mss_subtract   = 3
}