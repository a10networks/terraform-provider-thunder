provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ips_global_stats" "thunder_ips_global_stats" {

}
output "get_ips_global_stats" {
  value = ["${data.thunder_ips_global_stats.thunder_ips_global_stats}"]
}