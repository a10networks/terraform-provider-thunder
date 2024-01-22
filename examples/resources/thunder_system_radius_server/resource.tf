provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_radius_server" "thunder_system_radius_server" {
  accounting_interim_update = "ignore"
  accounting_on             = "ignore"
  accounting_start          = "append-entry"
  accounting_stop           = "delete-entry"
  attribute {
    attribute_value = "inside-ipv6-prefix"
    prefix_length   = "32"
    prefix_vendor   = 36007
    prefix_number   = 214
  }
  sampling_enable {
    counters1 = "all"
  }
}