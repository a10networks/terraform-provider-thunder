provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_token_authentication_players" "thunder_ddos_token_authentication_players" {
  dst_ip      = "10.10.10.14"
  dst_port    = 38148
  magic_value = 4046
  src_ip      = "10.10.10.10"
  src_port    = 43105
}