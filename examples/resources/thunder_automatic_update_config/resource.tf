provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_automatic_update_config" "thunder_automatic_update_config" {
  debug              = 1
  disable_ssl_verify = 1
  feature_name       = "ca-bundle"
  schedule           = 1
  weekly             = 1
  week_day           = "Wednesday"
  week_time          = "11:11"
}