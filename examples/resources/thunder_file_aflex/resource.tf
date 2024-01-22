provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_file_aflex" "thunder_file_aflex" {
  name          = "test_Aflex"
  protocol      = "http"
  host          = "10.64.3.190"
  path          = "/syslog.txt"
  use_mgmt_port = 1
}