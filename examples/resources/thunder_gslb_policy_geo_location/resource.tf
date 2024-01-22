provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_gslb_policy_geo_location" "thunder_gslb_policy_geo_location" {
  name        = "test"
  policy_name = "test"
  ip_multiple_fields {
    ip_sub      = "10.10.10.10"
    ip_mask_sub = "/24"
  }
  user_tag = "108"
}