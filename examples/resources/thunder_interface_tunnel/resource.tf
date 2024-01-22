provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_interface_tunnel" "thunder_interface_tunnel" {
  action = "enable"
  ifnum  = 2
  sampling_enable {
    counters1 = "all"
  }
}