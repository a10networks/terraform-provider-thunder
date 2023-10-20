provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ntp_auth_key" "resourceNtpAuthKeyTest" {
  key      = 1
  alg_type = "SHA1"
  key_type = "ascii"
  asc_key  = "a"
}