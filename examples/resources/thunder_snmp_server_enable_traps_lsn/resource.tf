
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_lsn"   "SnMPServerTrapsLsn"  {
       all = 1 
       total_port_usage_threshold = 0 
       per_ip_port_usage_threshold = 0 
       fixed_nat_port_mapping_file_change = 0 
       traffic_exceeded = 1
  
}
