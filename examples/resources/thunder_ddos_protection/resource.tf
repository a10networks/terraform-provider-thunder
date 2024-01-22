provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_protection" "thunder_ddos_protection" {
  toggle                       = "enable"
  rate_interval                = "1sec"
  force_routing_on_transp      = 1
  disable_on_reboot            = 1
  use_route                    = 1
  enable_now                   = 1
  mpls                         = 1
  src_dst_entry_limit          = "platform-default"
  src_zone_port_entry_limit    = "8M"
  non_zero_win_size_syncookie  = 1
  disallow_rst_ack_in_syn_auth = 1
}