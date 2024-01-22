provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_single_priority" "thunder_logging_single_priority" {
  levelname = "emergency"
}