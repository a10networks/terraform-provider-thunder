provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_techreport_interval" "thunder_techreport_interval" {
  value = 20
}