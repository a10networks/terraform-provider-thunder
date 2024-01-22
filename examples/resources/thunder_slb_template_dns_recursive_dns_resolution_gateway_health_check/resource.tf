provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dns_recursive_dns_resolution_gateway_health_check" "thunder_slb_template_dns_recursive_dns_resolution_gateway_health_check" {

  name                 = "test"
  gwhc_ns_cache_lookup = "disabled"
  interval             = 10
  num_query_type       = 42345
  query_name           = "a10networks.com"
  retry                = 6
  retry_multi          = 1
  timeout              = 5
  up_retry             = 1
}