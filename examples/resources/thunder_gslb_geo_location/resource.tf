provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_geo_location" "thunder_gslb_geo_location" {
  geo_locn_multiple_addresses {
    first_ipv6_address = "2001:db8:3333:4444:5555:6666:7777:8888"
    geol_ipv6_mask     = 122
  }
  geo_locn_obj_name = "test"
  user_tag          = "test"
}