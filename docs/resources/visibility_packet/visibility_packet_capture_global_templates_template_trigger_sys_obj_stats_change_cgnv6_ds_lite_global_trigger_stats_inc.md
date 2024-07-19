---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ds_lite_global_trigger_stats_inc Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ds_lite_global_trigger_stats_inc: Configure stats to trigger packet capture on increment
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ds_lite_global_trigger_stats_inc (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ds_lite_global_trigger_stats_inc`: Configure stats to trigger packet capture on increment

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ds_lite_global_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_ds_lite_global_trigger_stats_inc" {

  name                      = "test"
  fullcone_failure          = 1
  nat_port_unavailable_icmp = 1
  nat_port_unavailable_tcp  = 1
  nat_port_unavailable_udp  = 1
  user_quota_failure        = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `fullcone_failure` (Number) Enable automatic packet-capture for Full-cone Session Creation Failed
- `nat_port_unavailable_icmp` (Number) Enable automatic packet-capture for ICMP NAT Port Unavailable
- `nat_port_unavailable_tcp` (Number) Enable automatic packet-capture for TCP NAT Port Unavailable
- `nat_port_unavailable_udp` (Number) Enable automatic packet-capture for UDP NAT Port Unavailable
- `user_quota_failure` (Number) Enable automatic packet-capture for User-Quota Creation Failed
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

