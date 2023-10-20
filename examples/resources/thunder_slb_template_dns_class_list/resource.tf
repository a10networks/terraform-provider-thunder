provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_template_dns_class_list" "test_thunder_slb_template_dns_class_list" {
  name = "test_dns"
  lid_list {
    lidnum            = 1
    conn_rate_limit   = 200
    per               = 20
    over_limit_action = 1
    action_value      = "dns-cache-enable"
    lockout           = 30
    dns {
      cache_action              = "cache-enable"
      ttl                       = 20
      honor_server_response_ttl = 1
    }
    user_tag = "test_user"
  }
}
