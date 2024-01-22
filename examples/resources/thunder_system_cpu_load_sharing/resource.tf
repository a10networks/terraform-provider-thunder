provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_cpu_load_sharing" "thunder_system_cpu_load_sharing" {
  cpu_usage {
    low  = 60
    high = 75
  }
  disable = 1
  packets_per_second {
    min = 100000
  }
}