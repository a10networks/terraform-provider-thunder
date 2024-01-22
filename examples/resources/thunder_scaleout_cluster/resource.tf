provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster" "thunder_scaleout_cluster" {

  cluster_devices {
    enable = 0
    minimum_nodes {
      minimum_nodes_num = 0
    }
    cluster_discovery_timeout {
    }
    device_id_list {
      action = "enable"
    }
  }
  cluster_id = 43
  db_config {
    ticktime              = 6323
    initlimit             = 229
    synclimit             = 466
    minsessiontimeout     = 100
    maxsessiontimeout     = 30000
    client_recv_timeout   = 13000
    clientport            = 26310
    loopback_intf_support = 1
    broken_detect_timeout = 12000
    more_election_packet  = 1
    elect_conn_timeout    = 1200
  }
  device_groups {
    enable = 0
    device_group_list {
      device_group = 6
      device_id_list {
        device_id_start = device_id_start
        device_id_end   = device_id_end
      }
      user_tag = "55"
    }
  }
  local_device {
    priority     = 132
    id1          = 10
    action       = "enable"
    start_delay  = 103
    cluster_mode = "layer-2"
    l2_redirect {
      redirect_eth  = redirect_eth
      ethernet_vlan = 4005
      trunk_vlan    = 1244
    }
    traffic_redirection {
      follow_shared = 0
      interfaces {
        eth_cfg {
          ethernet = ethernet
        }
        trunk_cfg {
        }
        ve_cfg {
        }
        loopback_cfg {
        }
      }
      reachability_options {
        skip_default_route = 0
      }
    }
    session_sync {
      follow_shared = 0
      interfaces {
        eth_cfg {
          ethernet = ethernet
        }
        trunk_cfg {
        }
        ve_cfg {
        }
        loopback_cfg {
        }
      }
      reachability_options {
        skip_default_route = 0
      }
    }
    exclude_interfaces {
      eth_cfg {
        ethernet = ethernet
      }
      trunk_cfg {
      }
      ve_cfg {
      }
      loopback_cfg {
      }
    }
    tracking_template {
      template_list {
        template = "59"
        threshold_cfg {
          threshold = 39365
          action    = "down"
        }
        user_tag = "9"
      }
      multi_template_list {
        multi_template = "21"
        template {
          template_name = "58"
        }
        threshold = 43430
        action    = "down"
        user_tag  = "127"
      }
    }
  }
  service_config {
    enable = 0
    template_list {
      bucket_count = 256
      device_group = 8
      user_tag     = "78"
    }
  }
  slog_level = 5
  tracking_template {
    template_list {
      template = "13"
      threshold_cfg {
        threshold = 18851
        action    = "down"
      }
      user_tag = "24"
    }
  }

}