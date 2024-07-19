---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_authentication_global_trigger_stats_inc Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_authentication_global_trigger_stats_inc: Configure stats to trigger packet capture on increment
  PLACEHOLDER
---

# thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_authentication_global_trigger_stats_inc (Resource)

`thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_authentication_global_trigger_stats_inc`: Configure stats to trigger packet capture on increment

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_authentication_global_trigger_stats_inc" "thunder_visibility_packet_capture_global_templates_template_trigger_sys_obj_stats_change_aam_authentication_global_trigger_stats_inc" {

  name                     = "test"
  aflex_authz_fail         = 1
  authn_failure            = 1
  authz_failure            = 1
  connect_failed           = 1
  create_timer_failed      = 1
  dns_resolve_failed       = 1
  get_socket_option_failed = 1
  misses                   = 1
  open_socket_failed       = 1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `aflex_authz_fail` (Number) Enable automatic packet-capture for Total Authorization failure number in aFleX
- `authn_failure` (Number) Enable automatic packet-capture for Total Authentication failure number
- `authz_failure` (Number) Enable automatic packet-capture for Total Authorization failure number
- `connect_failed` (Number) Enable automatic packet-capture for Total AAM Connect Failed
- `create_timer_failed` (Number) Enable automatic packet-capture for Total AAM Timer Creation Failed
- `dns_resolve_failed` (Number) Enable automatic packet-capture for Total AAM DNS resolve failed
- `get_socket_option_failed` (Number) Enable automatic packet-capture for Total AAM Get Socket Option Failed
- `misses` (Number) Enable automatic packet-capture for Total Authentication Request Missed
- `open_socket_failed` (Number) Enable automatic packet-capture for Total AAM Open Socket Failed
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

