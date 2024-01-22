provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_flowspec" "thunder_flowspec" {
  name           = "test"
  src_addr_type  = "ip"
  src_ip_host    = "10.101.101.101"
  dest_addr_type = "ip"
  dest_ip_host   = "10.10.10.10"
  user_tag       = "test"
  source_port_list {
    port_attribute = "range"
    port_num       = 40
    port_num_end   = 44
  }


  destination_port_list {
    port_attribute = "range"
    port_num       = 200
    port_num_end   = 202
  }

  protocol_list {
    proto_attribute = "eq"
    proto_num       = 33
  }

  fragmentation_option_list {
    frag_attribute = "is-fragment"
  }


  dscp_list {
    dscp_attribute = "range"
    dscp_val       = 10
    dscp_val_end   = 20
  }


  filtering_action {
    redirect           = "next-hop-nlri"
    next_hop_nlri_type = "ip"
    ip_host_nlri       = "10.10.10.11"
    copy_ip_host_nlri  = 1
  }

}