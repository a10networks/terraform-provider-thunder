provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_mfa_management" "thunder_system_mfa_management" {
}