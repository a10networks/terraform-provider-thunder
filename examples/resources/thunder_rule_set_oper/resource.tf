provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_rule_set_oper" "thunder_rule_set_oper" {
  application {
    oper {
      protocol      = protocol
      rule_set_only = rule_set_only
      rule_list {
        stat_list {
          conns = conns
          bytes = bytes
        }
      }
    }
  }
  oper {
    policy_unmatched_drop = policy_unmatched_drop
    policy_permit         = policy_permit
    policy_deny           = policy_deny
    policy_reset          = policy_reset
    policy_rule_count     = policy_rule_count
    rule_stats {
      rule_hitcount = rule_hitcount
    }
    total_hit            = total_hit
    total_permit_bytes   = total_permit_bytes
    total_deny_bytes     = total_deny_bytes
    total_reset_bytes    = total_reset_bytes
    total_bytes          = total_bytes
    total_permit_packets = total_permit_packets
    total_deny_packets   = total_deny_packets
    total_reset_packets  = total_reset_packets
    total_packets        = total_packets
    total_active_tcp     = total_active_tcp
    total_active_udp     = total_active_udp
    total_active_icmp    = total_active_icmp
    total_active_others  = total_active_others
  }
  rule_list {
    oper {
      hitcount           = hitcount
      permitbytes        = permitbytes
      denybytes          = denybytes
      resetbytes         = resetbytes
      totalbytes         = totalbytes
      permitpackets      = permitpackets
      denypackets        = denypackets
      resetpackets       = resetpackets
      totalpackets       = totalpackets
      activesessiontcp   = activesessiontcp
      activesessionudp   = activesessionudp
      activesessionicmp  = activesessionicmp
      activesessionsctp  = activesessionsctp
      activesessionother = activesessionother
      activesessiontotal = activesessiontotal
      sessiontcp         = sessiontcp
      sessionudp         = sessionudp
      sessionicmp        = sessionicmp
      sessionsctp        = sessionsctp
      sessionother       = sessionother
      sessiontotal       = sessiontotal
      ratelimitdrops     = ratelimitdrops
    }
  }
  rules_by_zone {
    oper {
      group_list {
        rule_list {
          source_list {
          }
          dest_list {
          }
          service_list {
          }
          dscp_list {
          }
        }
      }
    }
  }
  track_app_rule_list {
    oper {
      rule_list {
      }
    }
  }

}
output "get_rule_set_oper" {
  value = ["${data.thunder_rule_set_oper.thunder_rule_set_oper}"]
}