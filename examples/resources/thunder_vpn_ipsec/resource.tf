provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vpn_ipsec" "thunder_vpn_ipsec" {
  name               = "test"
  anti_replay_window = "0"
  dh_group           = "0"
  dscp               = "default"

  enforce_traffic_selector = 0
  lifebytes                = 0
  lifetime                 = 28800
  mode                     = "tunnel"
  proto                    = "esp"
  sampling_enable {
    counters1 = "all"
  }
  sequence_number_disable = 0

  up       = 0
  user_tag = "68"
}