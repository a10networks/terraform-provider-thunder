provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_lif_bfd" "thunder_interface_lif_bfd" {

  ifname = "test"
  authentication {
    key_id   = 60
    method   = "md5"
    password = "12"
  }
  demand = 0
  echo   = 0
  interval_cfg {
    interval   = 431
    min_rx     = 683
    multiplier = 16
  }
}