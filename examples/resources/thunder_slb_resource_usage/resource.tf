provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_resource_usage" "Slb_Resource_Usage_Test" {
  slb_threshold_res_usage_percent = 60
  proxy_template_count            = 4096
  nat_pool_addr_count             = 2000
  fast_tcp_template_count         = 4096
  cache_template_count            = 2048
  health_monitor_count            = 1024
  fast_udp_template_count         = 4096
  virtual_port_count              = 8192
  client_ssl_template_count       = 8192
  persist_srcip_template_count    = 4096
  server_ssl_template_count       = 8192
  http_template_count             = 4096
  pbslb_subnet_count              = 65536
  persist_cookie_template_count   = 4096
  virtual_server_count            = 4096
  stream_template_count           = 4096
  conn_reuse_template_count       = 4096
  real_server_count               = 8192
  real_port_count                 = 16384
  service_group_count             = 8192
}