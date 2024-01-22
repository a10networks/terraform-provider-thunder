provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_accounting" "thunder_accounting" {

  commands = 0
  debug    = 12
  exec {
    accounting_exec_type   = "start-stop"
    accounting_exec_method = "tacplus"
  }
  stop_only = 0
  tacplus   = 0
}