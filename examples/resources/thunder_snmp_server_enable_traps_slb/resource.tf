
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_slb"   "SnMPServerEnableTrapsSlb"  {

    all= 1
    application_buffer_limit= 1
    gateway_up=1
    gateway_down= 0
    server_conn_limit= 0
    server_conn_resume= 1
    server_up= 1
    server_down= 1
    server_disabled= 0
    server_selection_failure= 0
    service_conn_limit= 0
    service_conn_resume= 0
    service_down= 1
    service_up= 0
    service_group_up= 0
    service_group_down= 0
    service_group_member_up= 0
    service_group_member_down= 0
    vip_connlimit= 0
    vip_connratelimit= 0
    vip_down= 0
    vip_port_connlimit= 0
    vip_port_connratelimit= 0
    vip_port_down= 0
    vip_port_up= 0
    vip_up= 0
    bw_rate_limit_exceed= 0
    bw_rate_limit_resume= 1
  
}
  
