provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_vpn" "thunder_vpn" {
  asymmetric_flow_support = 0
  crl {
  }
  default {
  }
  enable_vpn_metrics = 0
  error {
  }
  errordump {
  }
  extended_matching    = 0
  fragment_after_encap = 0
  group_list {
  }
  ike_acc_enable = 0
  ike_gateway_list {
    ike_version          = "v2"
    mode                 = "main"
    auth_method          = "preshare-key"
    preshare_key_value   = "69"
    hash                 = "sha256"
    interface_management = 0
    key                  = "132"
    key_passphrase       = "67"
    vrid {
      default = 0
    }
    local_cert {
      local_cert_name = "129"
    }
    remote_ca_cert {
      remote_cert_name = "65"
    }
    local_id  = "132"
    remote_id = "132"
    enc_cfg {
      encryption   = "des"
      hash         = "md5"
      prf          = "md5"
      priority     = 5
      gcm_priority = 5
    }
    dh_group = "1"
    local_address {
      local_ip = "10.10.10.10"
    }
    remote_address {
    }
    lifetime      = 86400
    fragment_size = 930
    nat_traversal = 0
    dpd {
      interval = 1805
      retry    = 7
    }
    disable_rekey         = 0
    configuration_payload = "dhcp"
    dhcp_server {
      pri {
        dhcp_pri_ipv4 = "10.10.10.10"
      }
      sec {
        dhcp_sec_ipv4 = "10.10.10.10"
      }
    }
    radius_server {
      radius_pri = "35"
      radius_sec = "35"
    }
    user_tag = "67"
    sampling_enable {
      counters1 = "all"
    }
  }
  ike_logging_enable = 0
  ike_sa {
  }
  ike_sa_brief {
  }
  ike_sa_clients {
  }
  ike_sa_timeout = 600
  ike_stats_by_gw {
  }
  ike_stats_global {
    sampling_enable {
      counters1 = "all"
    }
  }
  ipsec_cipher_check = 0
  ipsec_error_dump   = 0
  ipsec_group_list {
    ipsecgroup_cfg {
      priority = 10
    }
    user_tag = "69"
  }
  ipsec_list {
    mode     = "tunnel"
    dscp     = "default"
    proto    = "esp"
    dh_group = "0"
    enc_cfg {
      encryption   = "des"
      hash         = "md5"
      priority     = 5
      gcm_priority = 5
    }
    lifetime                = 28800
    lifebytes               = 0
    anti_replay_window      = "0"
    up                      = 0
    sequence_number_disable = 0
    traffic_selector {
      ipv4 {
        local_port           = 32771
        remote_ipv4_assigned = 0
        remote_port          = 32770
        protocol             = 131
      }
      ipv6 {
        localv6              = "2001:db8:3333:4444:5555:6666:7777:8888"
        local_portv6         = 32770
        remote_ipv6_assigned = 0
        remote_portv6        = 32768
        protocolv6           = 128
      }
    }
    enforce_traffic_selector = 0
    user_tag                 = "67"
    sampling_enable {
      counters1 = "all"
    }
    bind_tunnel {
      next_hop = "10.10.10.10"
    }
    ipsec_gateway {
      ike_gateway = "20"
    }
  }
  ipsec_mgmt_default_policy_drop = 0
  ipsec_sa {
  }
  ipsec_sa_by_gw {
  }
  ipsec_sa_clients {
  }
  ipsec_sa_stats_list {
    sampling_enable {
      counters1 = "all"
    }
  }
  log {
  }
  nat_traversal_flow_affinity = 0
  ocsp {
  }
  revocation_list {
    ca = "129"
    crl {
      crl_pri = "131"
      crl_sec = "129"
    }
    ocsp {
      ocsp_pri = "37"
      ocsp_sec = "37"
    }
    user_tag = "68"
  }
  sampling_enable {
    counters1 = "all"
  }
  signature_authentication = 0
  stateful_mode            = 0
  tcp_mss_adjust_disable   = 0
}