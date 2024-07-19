---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_object_templates_fw_server_port_tmpl_trigger_stats_inc Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_object_templates_fw_server_port_tmpl_trigger_stats_inc: Configure stats to triggers packet capture on increment
  PLACEHOLDER
---

# thunder_visibility_packet_capture_object_templates_fw_server_port_tmpl_trigger_stats_inc (Resource)

`thunder_visibility_packet_capture_object_templates_fw_server_port_tmpl_trigger_stats_inc`: Configure stats to triggers packet capture on increment

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_fw_server_port_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_fw_server_port_tmpl_trigger_stats_inc" {

  name                 = "test"
  es_resp_400          = 1
  es_resp_500          = 1
  es_resp_invalid_http = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `es_resp_400` (Number) Enable automatic packet-capture for Response status 400
- `es_resp_500` (Number) Enable automatic packet-capture for Response status 500
- `es_resp_invalid_http` (Number) Enable automatic packet-capture for Total non-http response
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

