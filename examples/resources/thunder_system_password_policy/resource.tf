provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_system_password_policy" "thunder_system_password_policy" {
  aging                        = "Strict"
  complexity                   = "Strict"
  forbid_consecutive_character = "0"
  history                      = "Strict"
  min_pswd_len                 = 8
  repeat_character_check       = "disable"
  username_check               = "disable"
}