provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vrrp_a_common" "thunder_vrrp_a_common" {
  device_id            = 3
  set_id               = 3
  disable_default_vrid = 1
  action               = "enable"
  hello_interval       = 3
  preemption_delay     = 222
  dead_timer           = 20
  arp_retry            = 3
  ttl                  = 120
  hop_limit            = 122
  track_event_delay    = 33
  get_ready_time       = 120
  inline_mode_cfg {
    inline_mode = 1
  }
  restart_time = 99
  hostid_append_to_vrid {
    hostid_append_to_vrid_default = 1
  }
  forward_l4_packet_on_standby = 1
}