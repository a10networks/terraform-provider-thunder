provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_route_map" "routemap" {
  action   = "permit"
  tag      = "map"
  sequence = 1
  user_tag = "routeMap"
  match {
    ip {
      address {
        prefix_list {
          name = "Prefix1"
        }
      }
      next_hop {
        acl1 = 50
      }
      peer {
        acl1 = 50
      }
      rib {
        reachable = "1.1.1.1"
      }
    }
    ipv6 {
      address_1 {
        prefix_list_2 {
          name = "Prefix1"
        }
      }
      next_hop_1 {
        next_hop_acl_name = "Ipv6Acl"
      }
      peer_1 {
        name = "ACL1"
      }
      rib {
        reachable = "1111::1"
      }
    }
    group {
      group_id = 2
      ha_state = "active"
    }
    as_path {
      name = "AS"
    }
    community {
      name_cfg {
        name        = "community"
        exact_match = 1
      }
    }
    extcommunity {
      extcommunity_l_name {
        name        = "exCommu"
        exact_match = 1
      }
    }
    scaleout {
      cluster_id        = 1
      operational_state = "up"
    }
    interface {
      loopback = 1
      ethernet = 0
      ve       = 0
      trunk    = 0
      tunnel   = 0
    }
    local_preference {
      val = 300
    }
    origin {
      igp        = 1
      egp        = 0
      incomplete = 0
    }
    metric {
      value = 200
    }
    route_type {
      external {
        value = "type-1"
      }
    }
    tag {
      value = 200
    }
  }
  set {
    ip {
      next_hop {
        address = "2.2.2.2"
      }
    }
    ipv6 {
      next_hop_1 {
        address = "3333::1"
        local {
          address = "4444::1"
        }
      }
    }
    level {
      value = "level-1"
    }
    metric {
      value = "100"
    }
    metric_type {
      value = "external"
    }
    tag {
      value = 2
    }
    aggregator {
      aggregator_as {
        ip  = "2.2.2.2"
        asn = 200
      }
    }
    as_path {
      prepend = "100 100 100"
    }
    atomic_aggregate = 1
    community        = "internet"
    local_preference {
      val = 300
    }
    origin {
      igp        = 1
      incomplete = 0
      egp        = 0
    }
    extcommunity {
      rt {
        value = "101:100"
      }
      soo {
        value = "101:101"
      }
    }
    originator_id {
      originator_ip = "1.1.1.1"
    }
    weight {
      weight_val = 200
    }
    dampening_cfg {
      dampening             = 1
      dampening_half_time   = 10
      dampening_max_supress = 60
      dampening_penalty     = 20
      dampening_reuse       = 1000
      dampening_supress     = 1200
    }
    comm_list {
      delete = 1
      v_std  = 50
    }
  }
}