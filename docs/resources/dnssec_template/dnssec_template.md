---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_dnssec_template Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_dnssec_template: template Settings
  PLACEHOLDER
---

# thunder_dnssec_template (Resource)

`thunder_dnssec_template`: template Settings

__PLACEHOLDER__

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `dnssec_temp_name` (String) DNSSEC Template Name

### Optional

- `algorithm` (String) 'RSASHA1': RSASHA1 algorithm; 'RSASHA256': RSASHA256 algorithm; 'RSASHA512': RSASHA512 algorithm;
- `combinations_limit` (Number) the max number of combinations per RRset (Default value is 31)
- `dnskey_ttl_k` (Number) The TTL value of DNSKEY RR
- `dnskey_ttl_v` (Number) in seconds, 14400 seconds by default
- `dnssec_template_ksk` (Block List, Max: 1) (see [below for nested schema](#nestedblock--dnssec_template_ksk))
- `dnssec_template_zsk` (Block List, Max: 1) (see [below for nested schema](#nestedblock--dnssec_template_zsk))
- `enable_nsec3` (Number) enable NSEC3 support. disabled by default
- `hsm` (String) specify the HSM template
- `return_nsec_on_failure` (Number) return NSEC/NSEC3 or not on failure case. return by default
- `signature_validity_period_k` (Number) The period that a signature is valid
- `signature_validity_period_v` (Number) in days, 10 days by default
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--dnssec_template_ksk"></a>
### Nested Schema for `dnssec_template_ksk`

Optional:

- `ksk_keysize_k` (Number) Specify the number of bits in the DNSSEC KSK keys
- `ksk_keysize_v` (Number) Default size is 2048 and must be an exact multiple of 64
- `ksk_lifetime_k` (Number) Set the lifetime for DNSSEC KSK keys in days
- `ksk_lifetime_v` (Number) Default value is 365 days
- `ksk_rollover_time_k` (Number) Set the rollover time in days
- `zsk_rollover_time_v` (Number) 7 days less than the lifetime by default


<a id="nestedblock--dnssec_template_zsk"></a>
### Nested Schema for `dnssec_template_zsk`

Optional:

- `zsk_keysize_k` (Number) Specify the number of bits in the DNSSEC ZSK keys
- `zsk_keysize_v` (Number) Default size is 2048 and must be an exact multiple of 64
- `zsk_lifetime_k` (Number) Set the lifetime for DNSSEC ZSK keys in days
- `zsk_lifetime_v` (Number) Default value is 90 days
- `zsk_rollover_time_k` (Number) Set the rollover time in days
- `zsk_rollover_time_v` (Number) 7 days less than the lifetime by default

