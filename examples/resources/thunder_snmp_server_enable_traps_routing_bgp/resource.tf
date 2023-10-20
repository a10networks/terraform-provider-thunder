
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_snmp_server_enable_traps_routing_bgp"   "SnMPServerEnableTrapsRoutingBgp"  {

         bgp_established_notification = 1 
         bgp_backward_trans_notification = 1  
         
}