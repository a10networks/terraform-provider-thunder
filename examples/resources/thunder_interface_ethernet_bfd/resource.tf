provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_ethernet_bfd" "thunder_interface_ethernet_bfd" {
  ifnum = 1
  authentication {
    key_id   = 255
    method   = "md5"
    password = "a10net"
  }
  demand = 0
  echo   = 0
  interval_cfg {
    interval   = 208
    min_rx     = 280
    multiplier = 47
  }
}