provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_lsn_rule_list_ip" "thunder_cgnv6_lsn_rule_list_ip" {

  name      = "test1"
  ipv4_addr = "192.168.0.101/32"
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "57"
}