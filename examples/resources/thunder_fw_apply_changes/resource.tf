provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_apply_changes" "test_thunder_fw_apply_changes" {
  forced = 0
}