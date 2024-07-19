---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_sip_trigger_stats_inc Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_sip_trigger_stats_inc: Configure stats to trigger packet capture on increment
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_sip_trigger_stats_inc (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_sip_trigger_stats_inc`: Configure stats to trigger packet capture on increment

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_sip_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_lsn_alg_sip_trigger_stats_inc" {

  name                  = "test"
  method_unknown        = 1
  parse_error           = 1
  tcp_out_of_order_drop = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `method_unknown` (Number) Enable automatic packet-capture for SIP Method UNKNOWN
- `parse_error` (Number) Enable automatic packet-capture for SIP Message Parse Error
- `tcp_out_of_order_drop` (Number) Enable automatic packet-capture for TCP Out-of-Order Drop
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

