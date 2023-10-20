provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_ip_tcp" "Iptcp" {
  syn_cookie {
    threshold = 32
  }
}