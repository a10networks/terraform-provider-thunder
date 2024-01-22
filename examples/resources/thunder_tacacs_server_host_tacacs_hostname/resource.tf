provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_tacacs_server_host_tacacs_hostname" "thunder_tacacs_server_host_tacacs_hostname" {
  hostname = "test"
  secret {
    source_eth   = 2
    secret_value = "test"
    port_cfg {
      port           = 22
      timeout        = 2
      monitor        = 1
      username       = "suraj"
      password       = 1
      password_value = "suraj"
    }
  }
}