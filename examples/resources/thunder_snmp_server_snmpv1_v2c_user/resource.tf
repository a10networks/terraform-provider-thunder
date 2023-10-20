provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}

resource "thunder_snmp_server_snmpv1_v2c_user" "Snmp_Server_SNMPv1_V2c_User_Test" {
  remote {
    host_list {
      dns_host = "testing"
    }
    ipv4_list {
      ipv4_host = "192.168.170.0"
      ipv4_mask = "255.255.255.0"
    }
    ipv6_list {
      ipv6_host = "2000:0db8:85a3:0000:0000:8a2e:0370:7336"
      ipv6_mask = 128
    }
  }
  passwd = "test"

  user_tag = "a"
  user     = "test_user"
  oid_list {
    remote {
      host_list {
        dns_host = "host_testing"

      }
    }

  }
}
