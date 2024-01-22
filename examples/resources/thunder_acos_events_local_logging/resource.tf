provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_local_logging" "thunder_acos_events_local_logging" {
  delete_old_logs_in_disk = 0
  enable                  = 0
  max_disk_space          = 100
  rate_limit              = 1000
  sampling_enable {
    counters1 = "all"
  }
}