provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vpn_ike_gateway" "thunder_vpn_ike_gateway" {
  name                  = "test"
  auth_method           = "preshare-key"
  configuration_payload = "dhcp"
  dhcp_server {
    pri {
      dhcp_pri_ipv4 = "10.10.10.10"
    }
    sec {
      dhcp_sec_ipv4 = "10.10.10.11"
    }
  }
  disable_rekey = 1
  dpd {
    interval = 1803
    retry    = 7
  }
  fragment_size        = 931
  ike_version          = "v2"
  interface_management = 0
  lifetime             = 86400
  local_address {
    local_ip = "10.10.10.10"
  }
  local_id      = "130"
  mode          = "main"
  nat_traversal = 0
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "66"
  vrid {
    default = 0
  }
}