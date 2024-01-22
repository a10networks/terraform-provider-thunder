provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_geo_location" "thunder_import_periodic_geo_location" {
  geo_location  = "test"
  period        = 6703081
  use_mgmt_port = 1
  remote_file   = "test"
}