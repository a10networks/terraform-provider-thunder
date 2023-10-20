

provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_gslb"   "SnMPServerEnableTrapsGslb"  {
    
       all = 1 
       zone = 1 
       site = 0 
       group = 1 
       service_ip = 1
     
} 
