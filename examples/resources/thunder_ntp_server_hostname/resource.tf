provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ntp_server_hostname" "thunder_ntp_server_hostname" {

  action          = "enable"
  host_servername = "5"
  key             = 41809
  prefer          = 0
}