provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ip_nat_pool_oper" "thunder_ip_nat_pool_oper" {
  oper {
    nat_pool_addr_list {
      port_usage  = port_usage
      total_used  = total_used
      total_freed = total_freed
      failed      = failed
    }
  }
  pool_name = "26"
}
output "get_ip_nat_pool_oper" {
  value = ["${data.thunder_ip_nat_pool_oper.thunder_ip_nat_pool_oper}"]
}