provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_nat_pool" "thunder_cgnv6_nat_pool" {
  end_address                            = "10.10.10.20"
  netmask                                = "/24"
  max_users_per_ip                       = 5
  partition                              = "test"
  per_batch_port_usage_warning_threshold = 34
  pool_name                              = "testPool"
  port_batch_v2_size                     = "64"
  shared                                 = 1
  simultaneous_batch_allocation          = 1
  start_address                          = "10.10.10.10"
  tcp_time_wait_interval                 = 5
  vrid                                   = 3
}