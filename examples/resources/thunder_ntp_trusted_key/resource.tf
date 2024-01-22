provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ntp_trusted_key" "thunder_ntp_trusted_key" {

  key = 47878
}