provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_email_buffer" "thunder_logging_email_buffer" {
  number = 100
  time   = 20
}