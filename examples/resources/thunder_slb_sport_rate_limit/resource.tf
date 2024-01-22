provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_sport_rate_limit" "thunder_slb_sport_rate_limit" {
  sampling_enable {
    counters1 = "all"
  }
}