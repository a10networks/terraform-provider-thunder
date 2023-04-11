---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_group Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_group: GSLB Group
  PLACEHOLDER
---

# thunder_gslb_group (Resource)

`thunder_gslb_group`: GSLB Group

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
    address  = var.dut9049
    username = var.username
    password = var.password
}

resource "thunder_gslb_group" "thunder_gslb_group_test"{
    name = "default"
    enable = 1
    priority = 255
    primary_list {
        primary = "20.121.5.226"
    }
    primary_ipv6_list {
        primary_ipv6 = "5001::1:20"
    }
    config_anywhere = 1
    config_merge = 1
    resolve_as = "resolve-to-ipv4"
    standalone = 1
    user_tag = "a10networks"
    auto_map_learn = 0
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Specify Group domain name

### Optional

- `auto_map_learn` (Number) IP Address learned from other controller
- `auto_map_primary` (Number) Primary Controller's IP address
- `auto_map_smart` (Number) Choose Best IP address
- `config_anywhere` (Number) Every member can do config
- `config_merge` (Number) Merge old master's config when new one take over
- `config_save` (Number) Accept config-save message from master
- `data_interface` (Number) Data Interface IP Address
- `dns_discover` (Number) Discover member via DNS Protocol
- `enable` (Number) Join GSLB Group
- `learn` (Number) Learn neighbour information from other controllers
- `mgmt_interface` (Number) Management Interface IP Address
- `primary_ipv6_list` (Block List) (see [below for nested schema](#nestedblock--primary_ipv6_list))
- `primary_list` (Block List) (see [below for nested schema](#nestedblock--primary_list))
- `priority` (Number) Specify Local Priority, default is 100
- `resolve_as` (String) 'resolve-to-ipv4': Use A Query only to resolve FQDN (Default Query type); 'resolve-to-ipv6': Use AAAA Query only to resolve FQDN; 'resolve-to-ipv4-and-ipv6': Use A as well as AAAA Query to resolve FQDN;
- `standalone` (Number) Run GSLB Group in standalone mode
- `suffix` (String) Set DNS Suffix (Name)
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--primary_ipv6_list"></a>
### Nested Schema for `primary_ipv6_list`

Optional:

- `primary_ipv6` (String) Specify Primary controller's IP address


<a id="nestedblock--primary_list"></a>
### Nested Schema for `primary_list`

Optional:

- `primary` (String) Specify Primary controller's IP address

