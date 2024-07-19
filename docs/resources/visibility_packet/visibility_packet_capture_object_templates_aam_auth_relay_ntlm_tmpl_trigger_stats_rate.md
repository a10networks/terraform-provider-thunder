---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_object_templates_aam_auth_relay_ntlm_tmpl_trigger_stats_rate Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_object_templates_aam_auth_relay_ntlm_tmpl_trigger_stats_rate: Configure stats to triggers packet capture on increment
  PLACEHOLDER
---

# thunder_visibility_packet_capture_object_templates_aam_auth_relay_ntlm_tmpl_trigger_stats_rate (Resource)

`thunder_visibility_packet_capture_object_templates_aam_auth_relay_ntlm_tmpl_trigger_stats_rate`: Configure stats to triggers packet capture on increment

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_relay_ntlm_tmpl_trigger_stats_rate" "thunder_visibility_packet_capture_object_templates_aam_auth_relay_ntlm_tmpl_trigger_stats_rate" {

  name                  = "test"
  buffer_alloc_fail     = 1
  duration              = 60
  encoding_fail         = 1
  failure               = 1
  insert_header_fail    = 1
  internal_error        = 1
  parse_header_fail     = 1
  threshold_exceeded_by = 5
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `buffer_alloc_fail` (Number) Enable automatic packet-capture for Buffer Allocation Failure
- `duration` (Number) Time in seconds to look for the anomaly, default is 60
- `encoding_fail` (Number) Enable automatic packet-capture for Encoding Failure
- `failure` (Number) Enable automatic packet-capture for Failure
- `insert_header_fail` (Number) Enable automatic packet-capture for Insert Header Failure
- `internal_error` (Number) Enable automatic packet-capture for Internal Error
- `parse_header_fail` (Number) Enable automatic packet-capture for Parse Header Failure
- `threshold_exceeded_by` (Number) Set the threshold to the number of times greater than the previous duration to start the capture, default is 5
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

