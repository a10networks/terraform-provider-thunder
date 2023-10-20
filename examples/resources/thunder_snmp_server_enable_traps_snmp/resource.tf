
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_snmp"   "SnMPServerEnableTrapsSnmp"  {

    all= 1
    linkdown= 0
    linkup= 1
   
}
