provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_urpf" "test_thunder_fw_urpf" {
  status = "strict"
}