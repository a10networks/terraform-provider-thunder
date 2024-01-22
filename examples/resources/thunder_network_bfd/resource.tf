provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_bfd" "thunder_network_bfd" {
  echo   = 1
  enable = 1
  interval_cfg {
    interval   = 800
    min_rx     = 800
    multiplier = 4
  }
}