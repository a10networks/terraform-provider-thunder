
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_vrrp_a"   "SnMPServerEnableTrapsVrrpA"  {

    all = 1
    active = 1
    standby = 1
   
}
  
