provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_timezone" "thunder_timezone" {
  timezone_index_cfg {
    timezone_index = "UTC"
    nodst          = 0
  }
}