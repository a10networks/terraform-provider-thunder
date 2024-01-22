provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_xml_schema" "thunder_import_periodic_xml_schema" {
  period        = 12415169
  xml_schema    = "43"
  use_mgmt_port = 1
  remote_file   = "test"
}