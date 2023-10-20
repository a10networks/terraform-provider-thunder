provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_interface_ethernet_bfd" "BFD" {
  ifnum = 5
  interval_cfg {
    interval   = 300
    min_rx     = 300
    multiplier = 3
  }
  echo   = 0
  demand = 0
  authentication {
    password  = "A10"
    encrypted = 1
    key_id    = 3
    method    = "md5"
  }

} 