provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_traffic_control_rule_set_rule" "thunder_traffic_control_rule_set_rule" {

  name          = "test"
  rule_set_name = "test"
  remark        = "test"
  status        = "enable"
  ip_version    = "v4"
  source_list {
    src_ip_subnet = "20.2.20.20/32"
  }

  src_threat_list = "test_threat"
  dest_list {
    dst_ip_subnet = "10.10.10.10/32"
  }

  dst_domain_list = "test_domain"
  dst_zone_any    = "any"
  dst_threat_list = "test_threat"
  service_list {
    protocols   = "tcp"
    gt_src_port = 20
    lt_dst_port = 24
  }
  user_tag = "test"
  sampling_enable {
    counters1 = "all"
  }

  action_group {
    limit_policy = 3
  }

}