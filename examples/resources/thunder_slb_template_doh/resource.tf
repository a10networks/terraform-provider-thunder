provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_doh" "thunder_slb_template_doh" {
  name       = "test1"
  conn_reuse = "disable"
  dns        = "default"
  dns_retry {
    retry_interval = 10
    after_timeout  = "close"
    max_trials     = 3
  }
  forwarder {
    v4_internal = 0
    v4_port     = 53
    v4_l4_proto = "both"
    v6_internal = 0
    v6_port     = 53
    v6_l4_proto = "both"
    bypass_doh  = 0
  }
  non_dns_request                     = "reject"
  reject_status_code                  = "400"
  shared_partition_tcp_proxy_template = 0
  user_tag                            = "77"
}