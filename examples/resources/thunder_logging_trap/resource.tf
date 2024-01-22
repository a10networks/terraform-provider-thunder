provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_logging_trap" "thunder_logging_trap" {
  trap_levelname = "alert"
}