provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_server" "thunder_slb_server" {
  name = "test-server1"
  host = "10.10.10.101"
  sampling_enable {
    counters1 = "all"
  }
}