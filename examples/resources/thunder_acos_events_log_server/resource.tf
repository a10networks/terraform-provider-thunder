provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_acos_events_log_server" "thunder_acos_events_log_server" {
  action = "enable"
  name   = "test"
  port_list {
    port_number = 346
    protocol    = "tcp"
    action      = "enable"
    user_tag    = "93"
    sampling_enable {
      counters1 = "all"
    }
  }
  resolve_as = "resolve-to-ipv4"
  sampling_enable {
    counters1 = "all"
  }
  server_ipv6_addr = "2001:db8:3333:4444:5555:6666:7777:8888"
  user_tag         = "127"
}