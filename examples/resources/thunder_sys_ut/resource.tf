provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sys_ut" "thunder_sys_ut" {
  action = "enable"
  state_list {
    name     = "a10"
    user_tag = "a11"
  }
}