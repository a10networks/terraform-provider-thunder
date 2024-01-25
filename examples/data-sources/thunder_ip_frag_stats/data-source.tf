provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ip_frag_stats" "thunder_ip_frag_stats" {
}
output "get_ip_frag_stats" {
  value = ["${data.thunder_ip_frag_stats.thunder_ip_frag_stats}"]
}