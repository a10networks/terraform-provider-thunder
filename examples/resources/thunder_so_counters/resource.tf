provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_so_counters" "thunder_so_counters" {
  sampling_enable {
    counters1 = "all"
  }
}