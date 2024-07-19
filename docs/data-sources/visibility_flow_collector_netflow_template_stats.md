---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_flow_collector_netflow_template_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_flow_collector_netflow_template_stats: Statistics for the object template
  PLACEHOLDER
---

# thunder_visibility_flow_collector_netflow_template_stats (Data Source)

`thunder_visibility_flow_collector_netflow_template_stats`: Statistics for the object template

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_flow_collector_netflow_template_stats" "thunder_visibility_flow_collector_netflow_template_stats" {
}
output "get_visibility_flow_collector_netflow_template_stats" {
  value = ["${data.thunder_visibility_flow_collector_netflow_template_stats.thunder_visibility_flow_collector_netflow_template_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `templates_added_to_delq` (Number) Netflow templates added to the delete queue
- `templates_removed_from_delq` (Number) Netflow templates removed from the delete queue

