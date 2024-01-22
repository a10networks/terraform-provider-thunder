provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_delete_geo_location" "thunder_delete_geo_location" {
  all = 0
}