provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_fixed_nat_global" "thunder_cgnv6_fixed_nat_global" {
  create_port_mapping_file = 0
  port_mapping_files_count = 5
  sampling_enable {
    counters1 = "all"
    counters2 = "nat44-udp-alg-fullcone-freed-shadow"
    counters3 = "fnat44_fwd_egress_pkt_size_range2"
    counters4 = "fnatdslite_fwd_ingress_pkt_size_range2"
  }
}