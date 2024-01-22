provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_techreport_max_logfile_size" "thunder_techreport_max_logfile_size" {
  value = 2
}