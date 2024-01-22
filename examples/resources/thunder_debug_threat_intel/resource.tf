provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_threat_intel" "thunder_debug_threat_intel" {
  dumy = 0
}