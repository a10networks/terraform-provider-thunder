provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_fw_tap_monitor" "test_thunder_fw_tap_monitor" {
  status = "enable"
  tap_port_cfg {
      tap_eth = 2
    }
}