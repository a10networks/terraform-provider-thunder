provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_nat_pool_group_member" "thunder_ipv6_nat_pool_group_member" {

  pool_name       = "11"
  pool_group_name = "35"


}