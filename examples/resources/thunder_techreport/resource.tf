provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_techreport" "thunder_techreport" {
  disable             = 1
  enable_full_history = 1
}