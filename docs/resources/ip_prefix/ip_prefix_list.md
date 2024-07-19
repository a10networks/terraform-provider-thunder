---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ip_prefix_list Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ip_prefix_list: Configure Prefix-list
  PLACEHOLDER
---

# thunder_ip_prefix_list (Resource)

`thunder_ip_prefix_list`: Configure Prefix-list

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_prefix_list" "IpPrefixList1" {
  name = "PREFIX-LIST"
  rules {
    le          = 32
    ge          = 24
    description = "Testing_resource"
    any         = 1
    action      = "deny"
    ipaddr      = "2.2.4.7/8"
    seq         = 1
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of a prefix list

### Optional

- `rules` (Block List) (see [below for nested schema](#nestedblock--rules))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--rules"></a>
### Nested Schema for `rules`

Optional:

- `action` (String) 'deny': Specify packets to reject; 'permit': Specify packets to forward;
- `any` (Number) Any prefix match. Same as "0.0.0.0/0 le 32"
- `description` (String) Prefix-list specific description (Up to 80 characters describing this prefix-list)
- `ge` (Number) Minimum prefix length to be matched
- `ipaddr` (String) IP prefix, e.g., 35.0.0.0/8
- `le` (Number) Maximum prefix length to be matched
- `seq` (Number) Sequence number of an entry

