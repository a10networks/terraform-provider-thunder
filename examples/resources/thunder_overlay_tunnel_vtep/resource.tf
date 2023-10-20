
provider  "thunder"  {
    address  = var.dut9049
    username = var.UserName
    password = var.Password
}
resource  "thunder_overlay_tunnel_vtep" "OverlayTunnelVtep"  {
  encap = "ip-encap"
	sampling_enable {
		counters1 = "all"
	}
	local_ip_address {
		     ip_address = "1.2.3.4"
		   }
	remote_ip_address_list {
	ip_address = "1.4.3.2"
	
	}
	id1 = 36
	}

