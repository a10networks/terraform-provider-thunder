provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_access_list" "thunder_ip_access_list" {
  name = "IPAccessList"
  rules {
    seq_num                  = 4
    action                   = "permit"
    icmp                     = 1
    any_type                 = 1
    src_any                  = 1
    dst_any                  = 1
    dscp                     = 48
    acl_log                  = 1
    transparent_session_only = 1
  }
}