provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_site_easy_rdt" "thunder_gslb_site_easy_rdt" {
  site_name     = "3"
  aging_time    = 10
  bind_geoloc   = 1
  ignore_count  = 11
  ipv6_mask     = 127
  limit         = 16383
  mask          = "/32"
  overlap       = 0
  range_factor  = 25
  smooth_factor = 12

}