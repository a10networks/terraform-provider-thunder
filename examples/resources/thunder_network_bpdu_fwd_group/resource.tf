provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_network_bpdu_fwd_group" "thunder_network_bpdu_fwd_group" {
  bpdu_fwd_group_number = 2
  ethernet_list {
    ethernet_start = 2
    ethernet_end   = 2
  }
  user_tag = "4"
}