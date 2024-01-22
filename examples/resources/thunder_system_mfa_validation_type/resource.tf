provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_mfa_validation_type" "thunder_system_mfa_validation_type" {
  ca_cert = "21"
}