provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_snmp_server_SNMPv1_v2c_user_oid" "thunder_snmp_server_SNMPv1_v2c_user_oid" {

  user    = "20"
  oid_val = "67"
  remote {
    ipv4_list {
      ipv4_host = "10.10.10.10"
    }
  }
  user_tag = "68"
}