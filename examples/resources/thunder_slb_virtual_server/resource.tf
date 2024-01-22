provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_virtual_server" "thunder_slb_virtual_server" {
  name       = "test"
  ip_address = "10.10.10.10"
}