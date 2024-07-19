---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ip_nat_alg_pptp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ip_nat_alg_pptp: PPTP ALG Settings
  PLACEHOLDER
---

# thunder_ip_nat_alg_pptp (Resource)

`thunder_ip_nat_alg_pptp`: PPTP ALG Settings

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_alg_pptp" "thunder_ip_nat_alg_pptp" {
  pptp = "enable"
  sampling_enable {
    counters1 = "all"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `pptp` (String) 'disable': Disable PPTP NAT ALG; 'enable': Enable PPTP NAT ALG;
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'current-smp-sessions': current-smp-sessions; 'current-gre-sessions': current-gre-sessions; 'smp-session-creation-failure': smp-session-creation-failure; 'truncated-pns-message': truncated-pns-message; 'truncated-pac-message': truncated-pac-message; 'mismatched-pns-call-id': mismatched-pns-call-id; 'mismatched-pac-call-id': mismatched-pac-call-id; 'retransmitted-pns-message': retransmitted-pns-message; 'retransmitted-pac-message': retransmitted-pac-message; 'truncated-gre-packet': truncated-gre-packet; 'unknown-gre-version': unknown-gre-version; 'no-matching-gre-session': no-matching-gre-session;

