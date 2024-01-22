provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_pool_group_member" "thunder_ip_nat_pool_group_member" {


  pool_group_name = "53"
  pool_name       = "a24"
}