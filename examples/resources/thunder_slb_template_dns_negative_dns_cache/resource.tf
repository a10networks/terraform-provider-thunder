provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dns_negative_dns_cache" "thunder_slb_template_dns_negative_dns_cache" {

  name                      = "test"
  bypass_query_threshold    = 3312
  enable_negative_dns_cache = 1
  max_negative_cache_ttl    = 24628
}