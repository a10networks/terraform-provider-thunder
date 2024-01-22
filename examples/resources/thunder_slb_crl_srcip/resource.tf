provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_crl_srcip" "test_thunder_slb_crl_srcip" {
  sampling_enable {
    counters1 = "all"
  }
}