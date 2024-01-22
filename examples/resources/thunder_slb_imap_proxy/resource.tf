provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_imap_proxy" "thunder_slb_imap_proxy" {
  sampling_enable {
    counters1 = "all"
  }
}