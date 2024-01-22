provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_mgmt_traffic" "thunder_ip_mgmt_traffic" {
  source_interface {
    loopback = 2
    ethernet = 2
  }
  traffic_type = "all"
}