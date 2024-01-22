provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_license_manager_overage" "thunder_license_manager_overage" {
  bytes   = 553
  days    = 134
  gb      = 57363
  hours   = 8
  kb      = 577
  mb      = 634
  minutes = 38
  seconds = 26
}