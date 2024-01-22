provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_service_group" "thunder_gslb_service_group" {
  dependency_site       = 1
  disable               = 1
  persistent_aging_time = 5
  persistent_ipv6_mask  = 128
  persistent_site       = 1
  service_group_name    = "sg80"
  user_tag              = "85"
}