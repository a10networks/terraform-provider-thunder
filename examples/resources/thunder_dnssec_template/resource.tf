provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_dnssec_template" "thunder_dnssec_template" {

  algorithm                   = "RSASHA1"
  combinations_limit          = 64547
  dnskey_ttl_k                = 0
  dnskey_ttl_v                = 14400
  dnssec_temp_name            = "test"
  enable_nsec3                = 0
  hsm                         = "test"
  return_nsec_on_failure      = 1
  signature_validity_period_k = 0
  signature_validity_period_v = 10
  user_tag                    = "test"
}