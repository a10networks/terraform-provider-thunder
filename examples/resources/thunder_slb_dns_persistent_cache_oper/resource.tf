provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_dns_persistent_cache_oper" "thunder_slb_dns_persistent_cache_oper" {
  oper {
    cache_entry {
      dnssec      = dnssec
      cache_type  = cache_type
      cache_class = cache_class
      q_length    = q_length
      r_length    = r_length
      ttl         = ttl
    }
    total_cache = total_cache
  }

}
output "get_slb_dns_persistent_cache_oper" {
  value = ["${data.thunder_slb_dns_persistent_cache_oper.thunder_slb_dns_persistent_cache_oper}"]
}