provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ip_address" "testname" {
  ip_addr = "3.3.3.3"
  ip_mask = "255.255.0.0"
}
