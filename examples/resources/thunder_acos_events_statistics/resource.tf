provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_statistics" "thunder_acos_events_statistics" {
  sampling_enable {
    counters1 = "all"
  }
}