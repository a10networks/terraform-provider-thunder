provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_site_active_rdt" "thunder_gslb_site_active_rdt" {
  site_name     = 3
  aging_time    = 123
  bind_geoloc   = 1
  ignore_count  = 11
  ipv6_mask     = 127
  limit         = 16383
  mask          = "/32"
  overlap       = 1
  range_factor  = 999
  smooth_factor = 15

}