provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dns_recursive_dns_resolution" "thunder_slb_template_dns_recursive_dns_resolution" {

  name                   = "test"
  csubnet_retry          = 1
  default_recursive      = 1
  dnssec_validation      = "disabled"
  fast_ns_selection      = "enabled"
  force_cname_resolution = "enabled"
  full_response          = 1
  gateway_health_check {
    gwhc_ns_cache_lookup = "disabled"
    interval             = 10
    num_query_type       = 42345
    query_name           = "a10networks.com"
    retry                = 6
    retry_multi          = 1
    timeout              = 5
    up_retry             = 1
  }
  lookup_order {
    query_type {
      num_query_type = 45027
      order          = "ipv4-precede-ipv6"
    }
  }
  max_trials                     = 255
  ns_cache_lookup                = "enabled"
  request_for_pending_resolution = "respond-with-servfail"
  retries_per_level              = 6
  udp_initial_interval           = 5
  udp_retry_interval             = 1
  use_client_qid                 = 0
  use_service_group_response     = "enabled"
}