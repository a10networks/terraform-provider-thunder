provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_email_level" "thunder_logging_email_level" {
  email_levelname = "alert"
}