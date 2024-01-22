provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_community_read" "thunder_snmp_server_community_read" {
  oid_list {
    oid_val = "7"
    remote {
      host_list {
        dns_host = "5"
      }
      ipv4_list {
        ipv4_host = "10.10.10.10"
      }
      ipv6_list {
        ipv6_host = "2001:db8:3333:4444:5555:6666:7777:8888"
        ipv6_mask = 21
      }
    }
    user_tag = "84"
  }
  remote {
    host_list {
      dns_host = "25"
    }
    ipv4_list {
      ipv4_host = "10.10.10.10"
    }
    ipv6_list {
      ipv6_host = "2001:db8:3333:4444:5555:6666:7777:8888"
      ipv6_mask = 52
    }
  }
  user     = "2"
  user_tag = "57"
}