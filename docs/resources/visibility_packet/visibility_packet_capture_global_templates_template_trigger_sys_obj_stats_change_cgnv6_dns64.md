---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64 Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64: Configure triggers for cgnv6.dns64 object
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64 (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64`: Configure triggers for cgnv6.dns64 object

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dns64" {

  name = "test"
  trigger_stats_inc {
    query_bad_pkt = 1
    resp_bad_pkt  = 1
    resp_bad_qr   = 1
    drop          = 1
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `trigger_stats_inc` (Block List, Max: 1) (see [below for nested schema](#nestedblock--trigger_stats_inc))
- `trigger_stats_rate` (Block List, Max: 1) (see [below for nested schema](#nestedblock--trigger_stats_rate))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--trigger_stats_inc"></a>
### Nested Schema for `trigger_stats_inc`

Optional:

- `drop` (Number) Enable automatic packet-capture for Dropped
- `query_bad_pkt` (Number) Enable automatic packet-capture for Query Bad Packet
- `resp_bad_pkt` (Number) Enable automatic packet-capture for Response Bad Packet
- `resp_bad_qr` (Number) Enable automatic packet-capture for Response Bad Query
- `uuid` (String) uuid of the object


<a id="nestedblock--trigger_stats_rate"></a>
### Nested Schema for `trigger_stats_rate`

Optional:

- `drop` (Number) Enable automatic packet-capture for Dropped
- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `query_bad_pkt` (Number) Enable automatic packet-capture for Query Bad Packet
- `resp_bad_pkt` (Number) Enable automatic packet-capture for Response Bad Packet
- `resp_bad_qr` (Number) Enable automatic packet-capture for Response Bad Query
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object

