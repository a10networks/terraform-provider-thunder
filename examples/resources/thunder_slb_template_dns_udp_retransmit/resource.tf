provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_dns_udp_retransmit" "thunder_slb_template_dns_udp_retransmit" {

  name           = "test"
  max_trials     = 4
  retry_interval = 11
}