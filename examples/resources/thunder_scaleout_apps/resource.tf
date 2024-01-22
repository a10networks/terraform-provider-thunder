provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_apps" "thunder_scaleout_apps" {

  enable = 1
  skip_mac_overwrite {
    enable = 1
  }
}