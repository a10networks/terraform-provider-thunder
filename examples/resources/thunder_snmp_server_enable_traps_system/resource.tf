
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_system"   "SnMPServerEnableTrapsSystem"  {
    all= 1
    control_cpu_high= 1
    data_cpu_high= 0
    fan= 0
    file_sys_read_only= 0
    high_disk_use= 0
    high_memory_use= 0
    high_temp= 1
    low_temp= 0
    license_management= 0
    packet_drop= 0
    power= 0
    pri_disk= 0
    restart= 0
    sec_disk= 0
    shutdown= 1
    smp_resource_event= 0
    syslog_severity_one= 0
    tacacs_server_up_down= 0
    start= 0
  
}
  
