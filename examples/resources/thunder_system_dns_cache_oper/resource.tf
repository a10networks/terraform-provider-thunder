provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_system_dns_cache_oper" "thunder_system_dns_cache_oper" {
}
output "get_system_dns_cache_oper" {
  value = ["${data.thunder_system_dns_cache_oper.thunder_system_dns_cache_oper}"]
}