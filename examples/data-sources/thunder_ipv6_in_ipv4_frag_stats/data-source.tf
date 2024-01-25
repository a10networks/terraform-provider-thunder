provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_ipv6_in_ipv4_frag_stats" "thunder_ipv6_in_ipv4_frag_stats" {
}
output "get_ipv6_in_ipv4_frag_stats" {
  value = ["${data.thunder_ipv6_in_ipv4_frag_stats.thunder_ipv6_in_ipv4_frag_stats}"]
}