provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_import_periodic_bw_list" "thunder_import_periodic_bw_list" {
  bw_list       = "test"
  period        = 267941
  use_mgmt_port = 1
  remote_file   = "test"
}