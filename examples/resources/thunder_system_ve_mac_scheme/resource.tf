provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_system_ve_mac_scheme" "resourceSystemVeMacSchemeTest" {
  ve_mac_scheme_val = "round-robin"
}