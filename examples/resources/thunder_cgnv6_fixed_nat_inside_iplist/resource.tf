provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_fixed_nat_inside_iplist" "thunder_cgnv6_fixed_nat_inside_iplist" {

  dest_rule_list    = "rule1"
  dynamic_pool_size = 20
  inside_ip_list    = "test"
  method            = "use-all-nat-ips"
  nat_ip_list       = "tett"
  offset {
    random = 1
  }
  partition           = "test"
  ports_per_user      = 2
  respond_to_user_mac = 1
  session_quota       = 22
  vrid                = 2

}