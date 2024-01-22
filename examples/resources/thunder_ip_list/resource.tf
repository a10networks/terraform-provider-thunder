provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_list" "thunder_ip_list" {
  name = "test"
  ipv4_config {
    ipv4_start_addr = "10.10.10.10"
    ipv4_end_addr   = "10.10.10.10"
  }
  user_tag = "83"
}