provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ntp_server_hostname" "NtpServerTest" {
  host_servername = "abcde"
  action          = "enable"
  prefer          = 0
}