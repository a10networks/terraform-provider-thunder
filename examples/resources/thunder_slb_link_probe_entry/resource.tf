provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_link_probe_entry" "thunder_slb_link_probe_entry" {
  sampling_enable {
    counters1 = "all"
  }
}