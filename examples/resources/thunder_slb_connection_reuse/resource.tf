provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_connection_reuse" "test_thunder_slb_connection_reuse" {
  sampling_enable {
    counters1 = "all"
  }
}