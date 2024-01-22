provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_hm_dplane" "thunder_slb_hm_dplane" {
  sampling_enable {
    counters1 = "all"
  }
}