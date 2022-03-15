---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_interface_ethernet_trunk_group Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_interface_ethernet_trunk_group (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **admin_key** (Number)
- **id** (String) The ID of this resource.
- **ifnum** (Number)
- **mode** (String)
- **port_priority** (Number)
- **timeout** (String)
- **trunk_number** (Number)
- **type** (String)
- **udld_timeout_cfg** (Block List, Max: 1) (see [below for nested schema](#nestedblock--udld_timeout_cfg))
- **user_tag** (String)
- **uuid** (String)

<a id="nestedblock--udld_timeout_cfg"></a>
### Nested Schema for `udld_timeout_cfg`

Optional:

- **fast** (Number)
- **slow** (Number)

