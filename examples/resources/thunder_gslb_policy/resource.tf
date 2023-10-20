provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_gslb_policy" "thunder_gslb_policy_test" {
  name                           = "a10"
  health_check_preference_enable = 1
  health_preference_top          = 1
  least_response                 = 1
  metric_fail_break              = 1
  metric_force_check             = 1
  metric_order                   = 0
  num_session_enable             = 1
  num_session_tolerance          = 17
  round_robin                    = 0
  weighted_alias                 = 1
  weighted_ip_enable             = 1
  weighted_ip_total_hits         = 1
  weighted_site_enable           = 1
  weighted_site_total_hits       = 1
  active_rdt {
    enable           = 1
    single_shot      = 1
    timeout          = 5
    skip             = 5
    keep_tracking    = 1
    samples          = 6
    tolerance        = 11
    difference       = 1
    limit            = 16382
    fail_break       = 1
    controller       = 1
    proto_rdt_enable = 1
  }
  active_servers_enable     = 1
  active_servers_fail_break = 1
  admin_ip_enable           = 1
  admin_ip_top_only         = 1
  admin_preference          = 1
  alias_admin_preference    = 1
  amount_first              = 1
  auto_map {
    ttl            = 200
    module_disable = 1
    all            = 1
  }
  bw_cost_enable     = 1
  bw_cost_fail_break = 1
  capacity {
    capacity_enable     = 1
    threshold           = 60
    capacity_fail_break = 1
  }
  connection_load {
    connection_load_enable     = 1
    connection_load_fail_break = 1
    connection_load_samples    = 1
    connection_load_interval   = 1
    limit                      = 1
    connection_load_limit      = 1
  }
  dns {
    server                = 1
    server_authoritative  = 1
    selected_only         = 1
    selected_only_value   = 1
    action                = 1
    active_only           = 1
    active_only_fail_safe = 1
    dns_auto_map          = 1
    backup_alias          = 1
    backup_server         = 1
    cache                 = 1
    aging_time            = 1000
  }
  edns {
    client_subnet_geographic = 1
  }
  geo_location_match {
    overlap          = 1
    geo_type_overlap = "global"
    match_first      = "policy"
  }
  geo_location_list {
    name = "GEO_LOCATION"
    ip_multiple_fields {
      ip_sub       = "10.23.16.129"
      ip_mask_sub  = "/23"
      ip_addr2_sub = "10.23.16.200"
    }
    ipv6_multiple_fields {
      ipv6_sub       = "4001::1:10"
      ipv6_mask_sub  = "64"
      ipv6_addr2_sub = "4001::1:100"
    }
    user_tag = "geo_loaction_tag"
  }
}
