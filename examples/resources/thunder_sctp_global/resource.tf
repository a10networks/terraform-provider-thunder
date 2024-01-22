provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_sctp_global" "thunder_sctp_global" {
  sampling_enable {
    counters1 = "all"
  }
}