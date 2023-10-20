
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_slb_change"   "SnMPServerEnableTrapsSlbChange"  {

    all= 1
    resource_usage_warning= 1
    connection_resource_event= 0
    server= 0
    server_port= 1
    ssl_cert_change= 0
    ssl_cert_expire= 0
    vip= 1
    vip_port= 0
    system_threshold= 1
  
}
