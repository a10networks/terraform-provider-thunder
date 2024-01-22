provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_alg_icmp" "test_thunder_fw_alg_icmp" {
  disable = "disable"
}