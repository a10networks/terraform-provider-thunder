provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_policy_geo_location_match" "thunder_gslb_policy_geo_location_match" {
  name             = "test"
  geo_type_overlap = "global"
  match_first      = "global"
  overlap          = 1
}