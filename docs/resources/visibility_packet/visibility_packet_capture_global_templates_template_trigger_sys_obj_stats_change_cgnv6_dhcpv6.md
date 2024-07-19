---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dhcpv6 Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dhcpv6: Configure triggers for cgnv6.dhcpv6.global object
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dhcpv6 (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dhcpv6`: Configure triggers for cgnv6.dhcpv6.global object

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dhcpv6" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_dhcpv6" {

  name = "test"
  trigger_stats_inc {
    packets_dropped           = 1
    pkts_dropped_during_clear = 1
    rcv_not_supported_msg     = 1
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

- `packets_dropped` (Number) Enable automatic packet-capture for Packets Dropped
- `pkts_dropped_during_clear` (Number) Enable automatic packet-capture for Packets dropped during clear bindings process
- `rcv_not_supported_msg` (Number) Enable automatic packet-capture for Packets with not supported messages sent
- `uuid` (String) uuid of the object


<a id="nestedblock--trigger_stats_rate"></a>
### Nested Schema for `trigger_stats_rate`

Optional:

- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `packets_dropped` (Number) Enable automatic packet-capture for Packets Dropped
- `pkts_dropped_during_clear` (Number) Enable automatic packet-capture for Packets dropped during clear bindings process
- `rcv_not_supported_msg` (Number) Enable automatic packet-capture for Packets with not supported messages sent
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object

