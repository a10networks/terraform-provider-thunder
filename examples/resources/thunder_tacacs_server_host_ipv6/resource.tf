provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_tacacs_server_host_ipv6" "thunder_tacacs_server_host_ipv6" {
  ipv6_addr = "2001:db8:3333:4444:5555:6666:7777:8888"
  secret {
    source_eth   = 2
    secret_value = "test"
    port_cfg {
      port           = 22
      timeout        = 11
      monitor        = 1
      username       = "suraj123"
      password       = 1
      password_value = "suraj"
    }
  }
}