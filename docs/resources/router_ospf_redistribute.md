---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_router_ospf_redistribute Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_router_ospf_redistribute (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **action** (String)
- **as_number** (String)
- **id** (String) The ID of this resource.
- **ip_nat** (Number)
- **ip_nat_floating_list** (Block List) (see [below for nested schema](#nestedblock--ip_nat_floating_list))
- **metric_ip_nat** (Number)
- **metric_type_ip_nat** (String)
- **ospf_list** (Block List) (see [below for nested schema](#nestedblock--ospf_list))
- **process_id** (Number)
- **redist_list** (Block List) (see [below for nested schema](#nestedblock--redist_list))
- **route_map_ip_nat** (String)
- **sequence** (String)
- **tag_ip_nat** (Number)
- **uuid** (String)
- **vip_floating_list** (Block List) (see [below for nested schema](#nestedblock--vip_floating_list))
- **vip_list** (Block List) (see [below for nested schema](#nestedblock--vip_list))

<a id="nestedblock--ip_nat_floating_list"></a>
### Nested Schema for `ip_nat_floating_list`

Optional:

- **ip_nat_floating_ip_forward** (String)
- **ip_nat_prefix** (String)


<a id="nestedblock--ospf_list"></a>
### Nested Schema for `ospf_list`

Optional:

- **metric_ospf** (Number)
- **metric_type_ospf** (String)
- **ospf** (Number)
- **process_id** (Number)
- **route_map_ospf** (String)
- **tag_ospf** (Number)


<a id="nestedblock--redist_list"></a>
### Nested Schema for `redist_list`

Optional:

- **metric** (Number)
- **metric_type** (String)
- **route_map** (String)
- **tag** (Number)
- **type** (String)


<a id="nestedblock--vip_floating_list"></a>
### Nested Schema for `vip_floating_list`

Optional:

- **vip_address** (String)
- **vip_floating_ip_forward** (String)


<a id="nestedblock--vip_list"></a>
### Nested Schema for `vip_list`

Optional:

- **metric_type_vip** (String)
- **metric_vip** (Number)
- **route_map_vip** (String)
- **tag_vip** (Number)
- **type_vip** (String)

