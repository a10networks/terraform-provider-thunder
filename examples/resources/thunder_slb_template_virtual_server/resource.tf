provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_virtual_server" "virtual-server" {
  name                            = "virtualserver1"
  conn_limit                      = 10
  conn_limit_reset                = 1
  conn_limit_no_logging           = 1
  subnet_gratuitous_arp           = 1
  conn_rate_limit                 = 300
  rate_interval                   = "100ms"
  conn_rate_limit_reset           = 1
  conn_rate_limit_no_logging      = 1
  disable_when_all_ports_down     = 1
  icmp_rate_limit                 = 3000
  icmp_lockup                     = 65535
  icmp_lockup_period              = 1600
  icmpv6_rate_limit               = 3001
  icmpv6_lockup                   = 60000
  icmpv6_lockup_period            = 16383
  tcp_stack_tfo_active_conn_limit = 1000
  tcp_stack_tfo_backoff_time      = 300
  tcp_stack_tfo_cookie_time_limit = 14400
  user_tag                        = "virtualserver1"
}