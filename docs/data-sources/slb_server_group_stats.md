---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_server_group_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_server_group_stats: Statistics for the object server-group
  PLACEHOLDER
---

# thunder_slb_server_group_stats (Data Source)

`thunder_slb_server_group_stats`: Statistics for the object server-group

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_server_group_stats" "thunder_slb_server_group_stats" {
  name = "server123"


}
output "get_slb_server_group_stats" {
  value = ["${data.thunder_slb_server_group_stats.thunder_slb_server_group_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) server-group name

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `dummy_conn` (Number) Current established connections

