provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_slb_player_id_global" "test_thunder_slb_player_id_global" {
  min_expiration          = 300
  pkt_activity_expiration = 234
  enable_64bit_player_id  = 1
  abs_max_expiration      = 200
  force_passive           = 1
  enforcement_timer       = 20
  sampling_enable {
    counters1 = "all"
  }
}

