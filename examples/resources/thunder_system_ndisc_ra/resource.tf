provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_ndisc_ra" "thunder_system_ndisc_ra" {
  sampling_enable {
    counters1 = "all"
  }
}