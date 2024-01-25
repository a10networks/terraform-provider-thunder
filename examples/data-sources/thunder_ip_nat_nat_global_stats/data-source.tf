provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ip_nat_nat_global_stats" "thunder_ip_nat_nat_global_stats" {
}
output "get_ip_nat_nat_global_stats" {
  value = ["${data.thunder_ip_nat_nat_global_stats.thunder_ip_nat_nat_global_stats}"]
}