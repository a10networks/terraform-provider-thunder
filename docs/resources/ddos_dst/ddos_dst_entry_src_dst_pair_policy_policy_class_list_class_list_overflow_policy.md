---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_class_list_overflow_policy Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_class_list_overflow_policy: Configure src dst dynamic entry count overflow policy for class-list
  PLACEHOLDER
---

# thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_class_list_overflow_policy (Resource)

`thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_class_list_overflow_policy`: Configure src dst dynamic entry count overflow policy for class-list

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_class_list_overflow_policy" "thunder_ddos_dst_entry_src_dst_pair_policy_policy_class_list_class_list_overflow_policy" {
  dst_entry_name        = "test"
  src_based_policy_name = "test"
  class_list_name       = "test"
  dummy_name            = "configuration"
  bypass                = 1
  exceed_log_cfg {
    log_enable = 1
  }
  log_periodic = 1
  glid         = "3"
  user_tag     = "test"
  l4_type_src_dst_overflow_list {
    protocol = "tcp"
    deny     = 1
    glid     = "3"
    user_tag = "test"
  }

  app_type_src_dst_overflow_list {
    protocol = "dns"
    user_tag = "test"
  }

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `class_list_name` (String) ClassListName
- `dst_entry_name` (String) DstEntryName
- `dummy_name` (String) 'configuration': Configure src dst dynamic entry count overflow policy for class-list;
- `src_based_policy_name` (String) SrcBasedPolicyName

### Optional

- `app_type_src_dst_overflow_list` (Block List) (see [below for nested schema](#nestedblock--app_type_src_dst_overflow_list))
- `bypass` (Number) Always permit for the Source to bypass all feature & limit checks
- `exceed_log_cfg` (Block List, Max: 1) (see [below for nested schema](#nestedblock--exceed_log_cfg))
- `glid` (String) Global limit ID
- `l4_type_src_dst_overflow_list` (Block List) (see [below for nested schema](#nestedblock--l4_type_src_dst_overflow_list))
- `log_periodic` (Number) Enable periodic log while event is continuing
- `template` (Block List, Max: 1) (see [below for nested schema](#nestedblock--template))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--app_type_src_dst_overflow_list"></a>
### Nested Schema for `app_type_src_dst_overflow_list`

Required:

- `protocol` (String) 'dns': dns; 'http': http; 'ssl-l4': ssl-l4; 'sip': sip;

Optional:

- `template` (Block List, Max: 1) (see [below for nested schema](#nestedblock--app_type_src_dst_overflow_list--template))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--app_type_src_dst_overflow_list--template"></a>
### Nested Schema for `app_type_src_dst_overflow_list.template`

Optional:

- `dns` (String) DDOS dns template
- `http` (String) DDOS http template
- `sip` (String) DDOS sip template
- `ssl_l4` (String) DDOS SSL-L4 template



<a id="nestedblock--exceed_log_cfg"></a>
### Nested Schema for `exceed_log_cfg`

Optional:

- `log_enable` (Number) Enable logging of limit exceed drop's


<a id="nestedblock--l4_type_src_dst_overflow_list"></a>
### Nested Schema for `l4_type_src_dst_overflow_list`

Required:

- `protocol` (String) 'tcp': tcp; 'udp': udp; 'icmp': icmp; 'other': other;

Optional:

- `deny` (Number) Blacklist and Drop all incoming packets for protocol
- `glid` (String) Global limit ID
- `template` (Block List, Max: 1) (see [below for nested schema](#nestedblock--l4_type_src_dst_overflow_list--template))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--l4_type_src_dst_overflow_list--template"></a>
### Nested Schema for `l4_type_src_dst_overflow_list.template`

Optional:

- `other` (String) DDOS OTHER template
- `tcp` (String) DDOS TCP template
- `template_icmp_v4` (String) DDOS icmp-v4 template
- `template_icmp_v6` (String) DDOS icmp-v6 template
- `udp` (String) DDOS UDP template



<a id="nestedblock--template"></a>
### Nested Schema for `template`

Optional:

- `logging` (String) DDOS logging template

