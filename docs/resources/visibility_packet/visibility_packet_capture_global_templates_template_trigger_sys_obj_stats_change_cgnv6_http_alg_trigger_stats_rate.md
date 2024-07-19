---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_http_alg_trigger_stats_rate Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_http_alg_trigger_stats_rate: Configure stats to trigger packet capture on increment rate
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_http_alg_trigger_stats_rate (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_http_alg_trigger_stats_rate`: Configure stats to trigger packet capture on increment rate

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_http_alg_trigger_stats_rate" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_cgnv6_http_alg_trigger_stats_rate" {

  name                     = "test"
  duration                 = 60
  header_insertion_failed  = 1
  header_removal_failed    = 1
  out_of_memory_dropped    = 1
  out_of_order_dropped     = 1
  queue_len_exceed_dropped = 1
  radius_requst_dropped    = 1
  radius_response_dropped  = 1
  threshold_exceeded_by    = 5
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `header_insertion_failed` (Number) Enable automatic packet-capture for Buff Insertion Failed
- `header_removal_failed` (Number) Enable automatic packet-capture for Buff Removal Failed
- `out_of_memory_dropped` (Number) Enable automatic packet-capture for Out-of-Memory Dropped
- `out_of_order_dropped` (Number) Enable automatic packet-capture for Packet Out-of-Order Dropped
- `queue_len_exceed_dropped` (Number) Enable automatic packet-capture for Queue Length Exceed Dropped
- `radius_requst_dropped` (Number) Enable automatic packet-capture for Query Request Dropped
- `radius_response_dropped` (Number) Enable automatic packet-capture for Query Response Dropped
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

