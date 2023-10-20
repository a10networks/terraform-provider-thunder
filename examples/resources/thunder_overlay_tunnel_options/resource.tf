
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_overlay_tunnel_options"   "OverlayTunnelOptions"  {
  
    ip_dscp_preserve =  1
    nvgre_disable_flow_id =  0
    nvgre_key_mode_lower24 =  1
    tcp_mss_adjust_disable =  1
    gateway_mac =  "2001:db8:0:200::1"
    vxlan_dest_port =  1
    
  }
  
