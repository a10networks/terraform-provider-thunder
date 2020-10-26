---
layout: "thunder"
page_title: "thunder: thunder_ip_nat_alg_pptp"
sidebar_current: "docs-thunder-resource-ip-nat-alg-pptp"
description: |-
  Provides details about thunder ip nat alg pptp resource for A10
---

# thunder\_ip\_nat\_alg\_pptp

`thunder_ip_nat_alg_pptp` Provides details about thunder ip nat alg pptp
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_ip_nat_alg_pptp" "pptp" {
  pptp = "disable"
    sampling_enable {
        counters1 = "all"
      }
}
```

## Argument Reference

* `pptp` - ‘disable’: Disable PPTP NAT ALG; ‘enable’: Enable PPTP NAT ALG;
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘current-smp-sessions’: current-smp-sessions; ‘current-gre-sessions’: current-gre-sessions; ‘smp-session-creation-failure’: smp-session-creation-failure; ‘truncated-pns-message’: truncated-pns-message; ‘truncated-pac-message’: truncated-pac-message; ‘mismatched-pns-call-id’: mismatched-pns-call-id; ‘mismatched-pac-call-id’: mismatched-pac-call-id; ‘retransmitted-pns-message’: retransmitted-pns-message; ‘retransmitted-pac-message’: retransmitted-pac-message; ‘truncated-gre-packet’: truncated-gre-packet; ‘unknown-gre-version’: unknown-gre-version; ‘no-matching-gre-session’: no-matching-gre-session;

