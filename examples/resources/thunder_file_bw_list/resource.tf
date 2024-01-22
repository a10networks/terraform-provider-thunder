provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_bw_list" "thunder_file_bw_list" {
  name          = "test_BWLIST"
  protocol      = "http"
  host          = "10.64.3.190"
  path          = "/bw-list.txt"
  use_mgmt_port = 1
}