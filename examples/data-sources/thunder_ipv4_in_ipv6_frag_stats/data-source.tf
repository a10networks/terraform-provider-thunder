provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ipv4_in_ipv6_frag_stats" "thunder_ipv4_in_ipv6_frag_stats" {
}
output "get_ipv4_in_ipv6_frag_stats" {
  value = ["${data.thunder_ipv4_in_ipv6_frag_stats.thunder_ipv4_in_ipv6_frag_stats}"]
}