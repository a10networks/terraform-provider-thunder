---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_dst_zone_port_zone_service_manual_mode Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_dst_zone_port_zone_service_manual_mode: Enter manual-mode configuration
  PLACEHOLDER
---

# thunder_ddos_dst_zone_port_zone_service_manual_mode (Resource)

`thunder_ddos_dst_zone_port_zone_service_manual_mode`: Enter manual-mode configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_zone_port_zone_service_manual_mode" "thunder_ddos_dst_zone_port_zone_service_manual_mode" {
  zone_name = "testZone"
  port_num  = 28712
  protocol  = "dns-tcp"
  config    = "configuration"
  user_tag  = "96"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `config` (String) 'configuration': Manual-mode configuration;
- `port_num` (String) PortNum
- `protocol` (String) Protocol
- `zone_name` (String) ZoneName

### Optional

- `close_sessions_for_unauth_sources` (Number) Close session for unauthenticated sources
- `glid_action` (String) 'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;
- `src_default_glid` (String) Global limit ID
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object
- `zone_template` (Block List, Max: 1) (see [below for nested schema](#nestedblock--zone_template))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--zone_template"></a>
### Nested Schema for `zone_template`

Optional:

- `dns` (String) DDOS dns template
- `encap` (String) DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)
- `http` (String) DDOS http template
- `quic` (String) DDOS quic template
- `sip` (String) DDOS sip template
- `ssl_l4` (String) DDOS ssl-l4 template
- `tcp` (String) DDOS tcp template
- `udp` (String) DDOS udp template

