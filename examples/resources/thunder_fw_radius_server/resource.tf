provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fw_radius_server" "thunder_fw_radius_server" {
  accounting_interim_update = "ignore"
  accounting_on             = "ignore"
  accounting_start          = "append-entry"
  accounting_stop           = "delete-entry"
  listen_port               = 38143
  sampling_enable {
    counters1 = "all"
  }
}