provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_session_filter" "thunder_session_filter" {
  filter_cfg {
    session_type = "ipv6"
    source_addr  = "10.10.10.10"
    source_port  = 699
    dest_addr    = "10.10.10.10"
    dport2       = 39947
  }
  set  = 1
  name = "test1"
}