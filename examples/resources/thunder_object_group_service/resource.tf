provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_object_group_service" "thunder_object_group_service" {

  svc_name    = "test"
  description = "60"
  rules {
    seq_num          = 3632
    protocol_id      = 17
    tcp_udp          = "tcp"
    icmp             = 0
    icmpv6           = 0
    icmp_type        = 235
    any_code         = 0
    icmpv6_type      = 44
    v6_any_code      = 0
    source           = 0
    eq_src           = 37103
    gt_src           = 2712
    lt_src           = 38948
    range_src        = 54829
    port_num_end_src = 50579
    eq_dst           = 48207
    gt_dst           = 61929
    lt_dst           = 15266
    range_dst        = 20743
    port_num_end_dst = 22510
    alg              = "FTP"
  }
  user_tag = "73"

}