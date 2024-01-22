provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_object_network" "thunder_object_network" {
  net_name       = "test"
  description    = "109"
  ip_range_end   = "10.10.10.101"
  ip_range_start = "10.10.10.10"
}