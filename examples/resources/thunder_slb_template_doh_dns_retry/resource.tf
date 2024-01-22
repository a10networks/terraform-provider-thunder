provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_doh_dns_retry" "thunder_slb_template_doh_dns_retry" {

  name           = "test1"
  after_timeout  = "close"
  max_trials     = 4
  retry_interval = 11
}