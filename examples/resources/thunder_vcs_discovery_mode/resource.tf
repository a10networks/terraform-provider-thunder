provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vcs_discovery_mode" "thunder_vcs_discovery_mode" {
  action = "Unicast"
}