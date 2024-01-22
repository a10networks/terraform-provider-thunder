provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_use_default_route" "thunder_ddos_use_default_route" {

  ethernet_start_cfg {
    ethernet_start = 1
    ethernet_end   = 2
  }

}