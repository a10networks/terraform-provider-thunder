provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_group" "thunder_gslb_group" {
  name     = "default"
  enable   = 1
  priority = 255
  primary_list {
    primary = "20.121.5.226"
  }
  primary_ipv6_list {
    primary_ipv6 = "5001::1:20"
  }
  config_anywhere = 1
  config_merge    = 1
  resolve_as      = "resolve-to-ipv4"
  standalone      = 1
  user_tag        = "a10networks"
  auto_map_learn  = 0
}