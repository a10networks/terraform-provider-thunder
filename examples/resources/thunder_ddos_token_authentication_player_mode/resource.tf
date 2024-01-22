provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_token_authentication_player_mode" "thunder_ddos_token_authentication_player_mode" {

  mode = "many-to-one"

}