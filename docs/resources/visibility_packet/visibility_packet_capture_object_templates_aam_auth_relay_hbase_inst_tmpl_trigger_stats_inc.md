---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_object_templates_aam_auth_relay_hbase_inst_tmpl_trigger_stats_inc Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_object_templates_aam_auth_relay_hbase_inst_tmpl_trigger_stats_inc: Configure stats to triggers packet capture on increment
  PLACEHOLDER
---

# thunder_visibility_packet_capture_object_templates_aam_auth_relay_hbase_inst_tmpl_trigger_stats_inc (Resource)

`thunder_visibility_packet_capture_object_templates_aam_auth_relay_hbase_inst_tmpl_trigger_stats_inc`: Configure stats to triggers packet capture on increment

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_object_templates_aam_auth_relay_hbase_inst_tmpl_trigger_stats_inc" "thunder_visibility_packet_capture_object_templates_aam_auth_relay_hbase_inst_tmpl_trigger_stats_inc" {

  name         = "test"
  bad_req      = 1
  forbidden    = 1
  no_creds     = 1
  not_found    = 1
  server_error = 1
  unauth       = 1
  unavailable  = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `bad_req` (Number) Enable automatic packet-capture for Bad Request
- `forbidden` (Number) Enable automatic packet-capture for Forbidden
- `no_creds` (Number) Enable automatic packet-capture for No Credential
- `not_found` (Number) Enable automatic packet-capture for Not Found
- `server_error` (Number) Enable automatic packet-capture for Internal Server Error
- `unauth` (Number) Enable automatic packet-capture for Unauthorized
- `unavailable` (Number) Enable automatic packet-capture for Service Unavailable
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

