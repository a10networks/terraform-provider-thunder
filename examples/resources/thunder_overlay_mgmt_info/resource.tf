provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_overlay_mgmt_info" "thunder_overlay_mgmt_info" {
  appstring   = "88"
  plugin_name = "4"
}