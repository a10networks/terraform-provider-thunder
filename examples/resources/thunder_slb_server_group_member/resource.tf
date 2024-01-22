provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_server_group_member" "thunder_slb_server_group_member" {
  server_group_name = "server_group"
  name              = "test"
}