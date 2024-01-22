provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_merge_mode_add_slb" "thunder_merge_mode_add_slb" {
  member              = 1
  server_port         = 1
  virtual_server_port = 1
}