provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_key_key" "thunder_key_key" {

  key_chain_flag = 1
  key_chain_name = "7"
  key_number     = 118
  key_string     = "61"
  user_tag       = "44"
}