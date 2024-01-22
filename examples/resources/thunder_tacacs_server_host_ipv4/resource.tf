provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_tacacs_server_host_ipv4" "thunder_tacacs_server_host_ipv4" {
  ipv4_addr = "10.10.10.10"
  secret {
    source_eth   = 2
    secret_value = "test"
    port_cfg {
      port           = 22
      timeout        = 3
      monitor        = 1
      username       = "pramod"
      password       = 1
      password_value = "pramod"
    }
  }
}