provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ntp_trusted_key" "resourceNtpTrustedKeyTest" {
  key = 2
}