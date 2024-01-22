provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_debug_monitor" "thunder_debug_monitor" {
  cpuid    = 2
  filename = "6"
  filesize = 1
  no_stop  = 1
  timeout  = 5
}