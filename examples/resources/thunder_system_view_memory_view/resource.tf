provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_view_memory_view" "thunder_system_view_memory_view" {
  sampling_enable {
    counters1 = "all"
  }
}