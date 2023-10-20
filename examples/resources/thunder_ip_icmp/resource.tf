provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ip_icmp" "ICMP" {
  redirect    = 1
  unreachable = 1
}