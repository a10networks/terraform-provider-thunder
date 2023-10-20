provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_logging_console" "A" {
  console_levelname = "debugging"
}