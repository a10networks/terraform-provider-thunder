provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_signature_extraction" "thunder_ddos_signature_extraction" {
  enable = 1
}