provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_tap" "thunder_ddos_tap" {
  ethernet_start_cfg {
    ethernet_start = 1
    ethernet_end   = 2
  }
}