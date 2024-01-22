provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_enable_traps_system" "thunder_snmp_server_enable_traps_system" {
  all = 0
  apps_global {
    sessions_threshold = 0
    cps_threshold      = 0
  }
  control_cpu_high      = 0
  data_cpu_high         = 0
  fan                   = 0
  file_sys_read_only    = 0
  high_disk_use         = 0
  high_memory_use       = 0
  high_temp             = 0
  license_management    = 0
  low_temp              = 0
  packet_drop           = 0
  power                 = 0
  pri_disk              = 0
  restart               = 0
  sec_disk              = 0
  shutdown              = 0
  smp_resource_event    = 0
  start                 = 0
  syslog_severity_one   = 0
  tacacs_server_up_down = 0
}