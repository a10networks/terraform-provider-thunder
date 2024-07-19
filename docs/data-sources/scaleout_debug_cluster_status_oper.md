---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_scaleout_debug_cluster_status_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_scaleout_debug_cluster_status_oper: Operational Status for the object status
  PLACEHOLDER
---

# thunder_scaleout_debug_cluster_status_oper (Data Source)

`thunder_scaleout_debug_cluster_status_oper`: Operational Status for the object status

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_scaleout_debug_cluster_status_oper" "thunder_scaleout_debug_cluster_status_oper" {
}
output "get_scaleout_debug_cluster_status_oper" {
  value = ["${data.thunder_scaleout_debug_cluster_status_oper.thunder_scaleout_debug_cluster_status_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `active_device_count` (Number)
- `active_device_count_list` (Block List) (see [below for nested schema](#nestedblock--oper--active_device_count_list))
- `cluster_disc_timer_running` (Number)
- `device_id` (Number)
- `explicitly_stop_service` (Number)
- `follow_vcs` (Number)
- `g_scaleout` (Number)
- `maintain_mode_configured` (Number)
- `min_node_num` (Number)
- `priority` (Number)
- `scaleout_current_role` (Number)
- `scaleout_device_group_list` (Block List) (see [below for nested schema](#nestedblock--oper--scaleout_device_group_list))
- `scaleout_device_id` (Number)
- `scaleout_device_list` (Block List) (see [below for nested schema](#nestedblock--oper--scaleout_device_list))
- `scaleout_enabled` (Number)
- `scaleout_priority` (Number)
- `vcs_device_id` (Number)
- `vcs_priority` (Number)

<a id="nestedblock--oper--active_device_count_list"></a>
### Nested Schema for `oper.active_device_count_list`

Optional:

- `device_id` (Number)
- `state` (Number)


<a id="nestedblock--oper--scaleout_device_group_list"></a>
### Nested Schema for `oper.scaleout_device_group_list`

Optional:

- `index` (Number)
- `members` (Number)
- `members_low64` (Number)


<a id="nestedblock--oper--scaleout_device_list"></a>
### Nested Schema for `oper.scaleout_device_list`

Optional:

- `device_id` (Number)
- `state` (Number)

