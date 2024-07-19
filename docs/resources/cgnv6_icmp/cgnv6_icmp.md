---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_icmp Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_icmp: CGNV6 ICMP Statistics
  PLACEHOLDER
---

# thunder_cgnv6_icmp (Resource)

`thunder_cgnv6_icmp`: CGNV6 ICMP Statistics

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_cgnv6_icmp" "thunder_cgnv6_icmp" {
  sampling_enable {
    counters1 = "all"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'icmp-unknown-type': ICMP Unknown Type; 'icmp-no-port-info': ICMP Port Info Not Included; 'icmp-no-session-drop': ICMP No Matching Session Drop; 'icmpv6-unknown-type': ICMPv6 Unknown Type; 'icmpv6-no-port-info': ICMPv6 Port Info Not Included; 'icmpv6-no-session-drop': ICMPv6 No Matching Session Drop; 'icmp-to-icmp': ICMP to ICMP Conversion; 'icmp-to-icmpv6': ICMP to ICMPv6 Conversion; 'icmpv6-to-icmp': ICMPv6 to ICMP Conversion; 'icmpv6-to-icmpv6': ICMPv6 to ICMPv6 Conversion; 'icmp-bad-type': Bad Embedded ICMP Type; 'icmpv6-bad-type': Bad Embedded ICMPv6 Type; '64-known-drop': NAT64 Forward Known ICMPv6 Drop; '64-unknown-drop': NAT64 Forward Unknown ICMPv6 Drop; '64-midpoint-hop': NAT64 Forward Unknown Source Drop; '46-known-drop': NAT64 Reverse Known ICMP Drop; '46-unknown-drop': NAT64 Reverse Known ICMPv6 Drop; '46-no-prefix-for-ipv4': NAT64 Reverse No Prefix Match for IPv4; '46-bad-encap-ip-header-len': 4to6 Bad Encapsulated IP Header Length; 'icmp-to-icmp-err': ICMP to ICMP Conversion Error; 'icmp-to-icmpv6-err': ICMP to ICMPv6 Conversion Error; 'icmpv6-to-icmp-err': ICMPv6 to ICMP Conversion Error; 'icmpv6-to-icmpv6-err': ICMPv6 to ICMPv6 Conversion Error; 'encap-cross-cpu-no-match': ICMP Embedded Cross CPU No Matching Session; 'encap-cross-cpu-preprocess-err': ICMP Embedded Cross CPU Preprocess Error; 'icmp-to-icmp-unknown-l4': ICMP Embedded Unknown L4 Protocol; 'icmp-to-icmpv6-unknown-l4': ICMP to ICMPv6 Embedded Unknown L4 Protocol; 'icmpv6-to-icmp-unknown-l4': ICMPv6 to ICMP Embedded Unknown L4 Protocol; 'icmpv6-to-icmpv6-unknown-l4': ICMPv6 to ICMPv6 Embedded Unknown L4 Protocol; 'static-nat': ICMP Static NAT; 'echo-to-pool-reply': Ping to Pool Reply; 'echo-to-pool-drop': Ping to Pool Drop; 'error-to-pool-drop': Error to Pool Drop; 'echo-to-pool-reply-v6': Ping6 to Pool Reply; 'echo-to-pool-drop-v6': Ping6 to Pool Drop; 'error-to-pool-drop-v6': Error to IPv6 Pool Drop; 'error-ip-mismatch': ICMP IP address mismatch;

