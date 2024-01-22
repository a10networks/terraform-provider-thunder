provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_mqtt" "thunder_slb_mqtt" {
  sampling_enable {
    counters1 = "all"
  }
}