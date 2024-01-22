provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_flowspec_filtering_action" "thunder_flowspec_filtering_action" {

  name               = "test"
  redirect           = "next-hop-nlri"
  next_hop_nlri_type = "ip"
  ip_host_nlri       = "10.10.10.11"
  copy_ip_host_nlri  = 1
}