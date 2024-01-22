provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_dnssec" "thunder_dnssec" {
  standalone = 1
  template_list {
    dnssec_temp_name            = "test"
    algorithm                   = "RSASHA1"
    combinations_limit          = 3
    dnskey_ttl_k                = 1
    dnskey_ttl_v                = 33
    enable_nsec3                = 1
    return_nsec_on_failure      = 1
    signature_validity_period_k = 1
    signature_validity_period_v = 11
    dnssec_template_zsk {
      zsk_lifetime_k      = 1
      zsk_lifetime_v      = 99
      zsk_rollover_time_k = 1
      zsk_rollover_time_v = 90
    }
    dnssec_template_ksk {
      ksk_lifetime_k      = 1
      ksk_lifetime_v      = 3
      ksk_rollover_time_k = 1
      zsk_rollover_time_v = 2
    }
    user_tag = "test"
  }
}