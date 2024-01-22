provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_email_address" "thunder_logging_email_address" {
  email_list {
    email_address = "21"
  }
}