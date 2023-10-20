provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_common_conn_rate_limit_src_ip" "Slb_Common_Conn_Rate_Limit_Src_Ip_Test" {
  protocol      = "tcp"
  log           = 1
  lock_out      = 1
  limit_period  = "100"
  limit         = 10000
  exceed_action = 1
  shared        = 1
  uuid          = "stringssss"
}
