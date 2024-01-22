provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vcs_stat" "thunder_vcs_stat" {
  sampling_enable {
    counters1 = "all"
  }
}