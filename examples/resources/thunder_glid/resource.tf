provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_glid" "thunder_glid" {
  name                     = 11
  conn_limit               = 464
  conn_rate_limit          = 20140
  conn_rate_limit_interval = 14827
  description              = "14"
  dns {
    action = "cache-disable"
    weight = 1
    ttl    = 16280
  }
  dns64 {
    disable          = 0
    exclusive_answer = 0
  }
  over_limit_cfg {
    over_limit_action = 1
    action_value      = "drop"
  }
  request_limit               = 270
  request_rate_limit          = 4655
  request_rate_limit_interval = 14807
  user_tag                    = "39"
}