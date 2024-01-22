provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_accounting_exec" "thunder_accounting_exec" {
  accounting_exec_method = "tacplus"
  accounting_exec_type   = "start-stop"
}