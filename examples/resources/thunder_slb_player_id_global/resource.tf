provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_player_id_global" "thunder_slb_player_id_global" {
  abs_max_expiration      = 10
  enable_64bit_player_id  = 0
  enforcement_timer       = 480
  force_passive           = 0
  min_expiration          = 1
  pkt_activity_expiration = 5
  sampling_enable {
    counters1 = "all"
  }
}