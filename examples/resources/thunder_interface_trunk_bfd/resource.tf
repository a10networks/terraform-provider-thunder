provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_trunk_bfd" "thunder_interface_trunk_bfd" {

  ifnum = 11
  authentication {
    key_id   = 195
    method   = "md5"
    password = "13"
  }
  demand = 0
  echo   = 0
  interval_cfg {
    interval   = 838
    min_rx     = 952
    multiplier = 14
  }
  per_member_port {
    local_address = "10.10.10.10"
  }
}