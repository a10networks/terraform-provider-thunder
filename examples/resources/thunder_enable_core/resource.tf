provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_enable_core" "thunder_enable_core" {
  core_level = "a10"
  full       = 1
}