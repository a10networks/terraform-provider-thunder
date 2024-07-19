---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_visibility_flow_collector_netflow_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_visibility_flow_collector_netflow_stats: Statistics for the object netflow
  PLACEHOLDER
---

# thunder_visibility_flow_collector_netflow_stats (Data Source)

`thunder_visibility_flow_collector_netflow_stats`: Statistics for the object netflow

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_visibility_flow_collector_netflow_stats" "thunder_visibility_flow_collector_netflow_stats" {
}
output "get_visibility_flow_collector_netflow_stats" {
  value = ["${data.thunder_visibility_flow_collector_netflow_stats.thunder_visibility_flow_collector_netflow_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))
- `template` (Block List, Max: 1) (see [below for nested schema](#nestedblock--template))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `agent_not_found` (Number) nflow pkts from not configured agents
- `frag_dropped` (Number) Total nflow fragment packets dropped
- `pkts_rcvd` (Number) Total nflow packets received
- `template_drop_exceeded` (Number) Total templates dropped because of maximum limit
- `template_drop_out_of_memory` (Number) Total templates dropped becuase of out of memory
- `unknown_dir` (Number) nflow sample direction is unknown
- `v10_templates_created` (Number) Total v10(IPFIX) templates created
- `v10_templates_deleted` (Number) Total v10(IPFIX) templates deleted
- `v9_templates_created` (Number) Total v9 templates created
- `v9_templates_deleted` (Number) Total v9 templates deleted
- `version_not_supported` (Number) nflow version not supported


<a id="nestedblock--template"></a>
### Nested Schema for `template`

Optional:

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--template--stats))

<a id="nestedblock--template--stats"></a>
### Nested Schema for `template.stats`

Optional:

- `templates_added_to_delq` (Number) Netflow templates added to the delete queue
- `templates_removed_from_delq` (Number) Netflow templates removed from the delete queue

