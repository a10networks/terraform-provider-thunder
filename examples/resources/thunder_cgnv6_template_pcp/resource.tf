provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_template_pcp" "thunder_cgnv6_template_pcp" {
  name                       = "test"
  allow_third_party_from_lan = 1
  allow_third_party_from_wan = 1
  announce                   = 1
  check_client_nonce         = 1
  disable_map_filter         = 1
  map                        = 1
  maximum                    = 1440
  minimum                    = 2
  pcp_server_port            = 5351
  peer                       = 1
  user_tag                   = "68"
}