provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_apps_skip_mac_overwrite" "thunder_scaleout_apps_skip_mac_overwrite" {
  enable = 1
}