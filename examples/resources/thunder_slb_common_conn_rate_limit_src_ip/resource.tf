provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_common_conn_rate_limit_src_ip" "Slb_Common_Conn_Rate_Limit_Src_Ip_Test" {
  disable_ipv6_support = 0
  protocol             = "tcp"
  lock_out             = 2425
  limit_period         = "100"
  limit                = 26381
  exceed_action        = 1
  shared               = 0
  log                  = 0
}