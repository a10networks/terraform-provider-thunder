provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_dns" "thunder_gslb_dns" {
  action  = "none"
  logging = "none"
  sampling_enable {
    counters1 = "all"
  }
}