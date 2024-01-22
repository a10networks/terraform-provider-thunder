provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_fail_safe" "thunder_fail_safe" {
  dataplane_recovery_timeout = 5
  disable_failsafe {
    action = "all"
  }
  fpga_buff_recovery_threshold   = 2
  fpga_monitor_interval          = 1
  fpga_monitor_threshold         = 30
  hw_error_monitor               = "hw-error-monitor-enable"
  hw_error_recovery_timeout      = 336
  hw_ssl_timeout_monitor         = "hw-ssl-timeout-monitor-enable"
  kill                           = 1
  session_mem_recovery_threshold = 30
  sw_error_monitor_enable        = 1
  sw_error_recovery_timeout      = 3
  total_memory_size_check        = 112
}