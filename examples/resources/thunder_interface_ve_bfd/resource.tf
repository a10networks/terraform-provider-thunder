provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_interface_ve_bfd" "VEbfd" {
  ifnum = 300
  interval_cfg {
    interval   = 100
    min_rx     = 100
    multiplier = 3
  }
  demand = 0
  echo   = 0
  authentication {
    password = "bfdis"
    key_id   = 5
    method   = "md5"
  }
}