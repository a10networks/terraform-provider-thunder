provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dns" "t_dns_01" {

  name = "t_dns_01"
  class_list {
    name = "CLIST_A"
    lid_list {
      lidnum            = 10
      conn_rate_limit   = 3000
      per               = 10
      over_limit_action = 1
      action_value      = "dns-cache-enable"
      lockout           = 30
      log               = 0
      dns {
        cache_action              = "cache-disable"
        honor_server_response_ttl = 1
      }
    }
  }
}