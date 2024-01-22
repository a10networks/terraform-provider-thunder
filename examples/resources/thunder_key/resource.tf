provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_key" "thunder_key" {
  key_chain_flag = 1
  key_chain_name = "7"
  key_list {
    key_number = 39
    key_string = "112"
    user_tag   = "88"
  }
  user_tag = "20"
}