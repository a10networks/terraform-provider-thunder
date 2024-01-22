provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_axdebug_pcapng_config" "thunder_axdebug_pcapng_config" {
  exit           = 0
  pcapng_enable  = 0
  ssl_key_enable = 0
}