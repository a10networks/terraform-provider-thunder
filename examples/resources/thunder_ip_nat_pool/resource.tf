provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_pool" "test12"{
  pool_name = "test"
      start_address = "10.10.10.20"
      end_address = "10.10.10.22"
      netmask = "/24"
      ip_rr = 1
      port_overload = 1

}
