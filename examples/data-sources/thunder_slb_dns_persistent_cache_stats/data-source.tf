provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_dns_persistent_cache_stats" "thunder_slb_dns_persistent_cache_stats" {

}
output "get_slb_dns_persistent_cache_stats" {
  value = ["${data.thunder_slb_dns_persistent_cache_stats.thunder_slb_dns_persistent_cache_stats}"]
}