provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_one_to_one_pool" "thunder_cgnv6_one_to_one_pool" {
  end_address   = "10.10.10.20"
  netmask       = "/24"
  partition     = "test"
  pool_name     = "52"
  shared        = 1
  start_address = "10.10.10.10"
  vrid          = 3
}