provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vcs_unicast_election_members" "thunder_vcs_unicast_election_members" {
  ip_address_cfg {
    ip_address    = "10.10.10.10"
    use_mgmt_port = 1
  }
}