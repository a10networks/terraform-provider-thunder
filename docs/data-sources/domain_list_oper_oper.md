---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_domain_list_oper_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_domain_list_oper_oper: Operational Status for the object domain-list-oper
  PLACEHOLDER
---

# thunder_domain_list_oper_oper (Data Source)

`thunder_domain_list_oper_oper`: Operational Status for the object domain-list-oper

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_domain_list_oper_oper" "thunder_domain_list_oper_oper" {

}
output "get_domain_list_oper_oper" {
  value = ["${data.thunder_domain_list_oper_oper.thunder_domain_list_oper_oper}"]
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

- `binding_info` (Number)
- `displayed_count` (Number)
- `domain_entries` (Block List) (see [below for nested schema](#nestedblock--oper--domain_entries))
- `domain_list` (Block List) (see [below for nested schema](#nestedblock--oper--domain_list))
- `domain_list_name_filter` (String)
- `domain_list_type` (String)
- `match_type` (String)

<a id="nestedblock--oper--domain_entries"></a>
### Nested Schema for `oper.domain_entries`

Optional:

- `domain_match_type` (String)
- `domain_name` (String)
- `zone_transfer_interval` (String)
- `zone_transfer_ip` (String)
- `zone_transfer_port` (String)


<a id="nestedblock--oper--domain_list"></a>
### Nested Schema for `oper.domain_list`

Optional:

- `binding_info_list` (Block List) (see [below for nested schema](#nestedblock--oper--domain_list--binding_info_list))
- `binding_num` (Number)
- `config_type` (String)
- `domain_list_name` (String)
- `total_entry_num` (Number)

<a id="nestedblock--oper--domain_list--binding_info_list"></a>
### Nested Schema for `oper.domain_list.binding_info_list`

Optional:

- `domain_group_name` (String)

