provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ipv6_access_list" "thunder_ipv6_access_list" {
  name = "s11"
  rules {
    seq_num   = 7397
    action    = "deny"
    remark    = "24"
    icmp      = 1
    src_any   = 1
    dst_any   = 1
    fragments = 1
    ethernet  = 2
    dscp      = 11
    acl_log   = 1
  }
  user_tag = "86"
}