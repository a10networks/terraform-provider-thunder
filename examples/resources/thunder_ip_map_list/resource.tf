provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_map_list" "thunder_ip_map_list" {
  name = "test"
  file = 0
}