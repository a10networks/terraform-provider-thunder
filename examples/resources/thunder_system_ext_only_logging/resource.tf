provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_ext_only_logging" "thunder_system_ext_only_logging" {
}