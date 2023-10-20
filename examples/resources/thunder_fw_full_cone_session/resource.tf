provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_full_cone_session" "test_thunder_fw_full_cone_session" {
  uuid = "string"
}