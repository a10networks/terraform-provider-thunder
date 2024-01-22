provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_io_cpu" "thunder_system_io_cpu" {
  max_cores = 2
}