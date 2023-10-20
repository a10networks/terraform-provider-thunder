provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_timezone" "timezoneTest" {
  timezone_index_cfg {
    timezone_index = "America/Phoenix"
    nodst          = 1
  }
}
