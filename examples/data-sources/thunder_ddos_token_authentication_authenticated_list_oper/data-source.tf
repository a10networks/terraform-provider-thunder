provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ddos_token_authentication_authenticated_list_oper" "thunder_ddos_token_authentication_authenticated_list_oper" {
  oper {
    player_list {
      src_port = src_port
      dst_port = dst_port
      token    = token
    }
  }
}
output "get_ddos_token_authentication_authenticated_list_oper" {
  value = ["${data.thunder_ddos_token_authentication_authenticated_list_oper.thunder_ddos_token_authentication_authenticated_list_oper}"]
}