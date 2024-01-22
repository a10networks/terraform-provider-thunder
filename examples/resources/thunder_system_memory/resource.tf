provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_memory" "thunder_system_memory" {
  sampling_enable {
    counters1 = "all"
  }
}