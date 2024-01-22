provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_SNMPv1_v2c_user" "thunder_snmp_server_SNMPv1_v2c_user" {
  oid_list {
    oid_val = "67"
    remote {
      ipv4_list {
        ipv4_host = "10.11.22.10"
      }
    }
  }
  user     = "20"
  user_tag = "65"
}