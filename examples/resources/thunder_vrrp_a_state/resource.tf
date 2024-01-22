provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_state" "thunder_vrrp_a_state" {
  sampling_enable {
    counters1 = "all"
  }
}