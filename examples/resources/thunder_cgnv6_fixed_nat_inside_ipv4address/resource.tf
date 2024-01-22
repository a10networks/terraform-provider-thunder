provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_fixed_nat_inside_ipv4address" "thunder_cgnv6_fixed_nat_inside_ipv4address" {

  dynamic_pool_size    = 1
  inside_end_address   = "25.25.25.25"
  inside_netmask       = "/32"
  inside_start_address = "25.25.25.25"
  method               = "use-all-nat-ips"
  nat_end_address      = "30.30.30.80"
  nat_netmask          = "/24"
  nat_start_address    = "30.30.30.80"
  offset {
    random = 1
  }
  partition           = "test"
  ports_per_user      = 22
  respond_to_user_mac = 1
  session_quota       = 200
  vrid                = 2

}