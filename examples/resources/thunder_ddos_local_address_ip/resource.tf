provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_local_address_ip" "thunder_ddos_local_address_ip" {
  ip_addr = "10.10.10.10"
}