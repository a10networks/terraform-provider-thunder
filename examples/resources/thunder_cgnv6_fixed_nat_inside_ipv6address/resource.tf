provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_fixed_nat_inside_ipv6address" "thunder_cgnv6_fixed_nat_inside_ipv6address" {
  dynamic_pool_size    = 22
  inside_end_address   = "2001:db8::2:11"
  inside_netmask       = 128
  inside_start_address = "2001:db8::2:11"
  method               = "use-all-nat-ips"
  nat_end_address      = "30.30.30.90"
  nat_netmask          = "/24"
  nat_start_address    = "30.30.30.90"
  offset {
    random = 1
  }
  partition           = "test"
  ports_per_user      = 23
  respond_to_user_mac = 1
  session_quota       = 100

}