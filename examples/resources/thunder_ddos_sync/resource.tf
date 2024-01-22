provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_sync" "thunder_ddos_sync" {
  enable = 1
  peer_ip_cfg {
    peer_ip = "10.10.10.10"
  }
}