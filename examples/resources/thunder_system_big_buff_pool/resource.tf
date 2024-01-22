provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_big_buff_pool" "thunder_system_big_buff_pool" {
  big_buff_pool = 0
}