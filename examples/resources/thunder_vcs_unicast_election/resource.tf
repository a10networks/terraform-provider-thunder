provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vcs_unicast_election" "thunder_vcs_unicast_election" {
  members {
    ip_address_cfg {
      ip_address = "10.10.10.10"
    }
  }
  port = 41217
}