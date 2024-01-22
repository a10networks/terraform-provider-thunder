provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_player_id_list" "thunder_slb_player_id_list" {
  player_record {
    player_id           = 2097859606
    game_server_ipv4    = "10.10.10.10"
    game_server_ipv6    = "2001:db8:3333:4444:5555:6666:7777:8888"
    game_server_port_v4 = 19381
    game_server_port_v6 = 11622
  }
}