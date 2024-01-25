provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_player_id_ep_oper" "thunder_slb_player_id_ep_oper" {
}
output "get_slb_player_id_ep_oper" {
  value = ["${data.thunder_slb_player_id_ep_oper.thunder_slb_player_id_ep_oper}"]
}