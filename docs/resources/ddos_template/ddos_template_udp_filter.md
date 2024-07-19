---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_template_udp_filter Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_template_udp_filter: UDP Filter Configuration
  PLACEHOLDER
---

# thunder_ddos_template_udp_filter (Resource)

`thunder_ddos_template_udp_filter`: UDP Filter Configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_template_udp_filter" "thunder_ddos_template_udp_filter" {
  name                 = "test"
  udp_filter_action    = "blacklist-src"
  udp_filter_regex     = "1252"
  udp_filter_seq       = 3
  udp_filter_unmatched = 1
  user_tag             = "94"

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name
- `udp_filter_seq` (Number) Sequence number

### Optional

- `byte_offset_filter` (String) Filter Expression using Berkeley Packet Filter syntax
- `udp_filter_action` (String) 'blacklist-src': Also blacklist the source when action is taken; 'whitelist-src': Whitelist the source after filter passes, packets are dropped until then; 'count-only': Take no action and continue processing the next filter;
- `udp_filter_regex` (String) Regex Expression
- `udp_filter_unmatched` (Number) action taken when it does not match
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

