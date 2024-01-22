provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_call_home_profile" "thunder_call_home_profile" {
  action = "register"
  time   = 10
}