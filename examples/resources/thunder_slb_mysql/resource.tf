provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_mysql" "thunder_slb_mysql" {
  sampling_enable {
    counters1 = "all"
  }
}