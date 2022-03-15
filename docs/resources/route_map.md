---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_route_map Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  
---

# thunder_route_map (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **action** (String)
- **id** (String) The ID of this resource.
- **match** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match))
- **sequence** (Number)
- **set** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set))
- **tag** (String)
- **user_tag** (String)
- **uuid** (String)

<a id="nestedblock--match"></a>
### Nested Schema for `match`

Optional:

- **as_path** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--as_path))
- **community** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--community))
- **extcommunity** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--extcommunity))
- **group** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--group))
- **interface** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--interface))
- **ip** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ip))
- **ipv6** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ipv6))
- **local_preference** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--local_preference))
- **metric** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--metric))
- **origin** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--origin))
- **route_type** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--route_type))
- **scaleout** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--scaleout))
- **tag** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--tag))
- **uuid** (String)

<a id="nestedblock--match--as_path"></a>
### Nested Schema for `match.as_path`

Optional:

- **name** (String)


<a id="nestedblock--match--community"></a>
### Nested Schema for `match.community`

Optional:

- **name_cfg** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--community--name_cfg))

<a id="nestedblock--match--community--name_cfg"></a>
### Nested Schema for `match.community.name_cfg`

Optional:

- **exact_match** (Number)
- **name** (String)



<a id="nestedblock--match--extcommunity"></a>
### Nested Schema for `match.extcommunity`

Optional:

- **extcommunity_l_name** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--extcommunity--extcommunity_l_name))

<a id="nestedblock--match--extcommunity--extcommunity_l_name"></a>
### Nested Schema for `match.extcommunity.extcommunity_l_name`

Optional:

- **exact_match** (Number)
- **name** (String)



<a id="nestedblock--match--group"></a>
### Nested Schema for `match.group`

Optional:

- **group_id** (Number)
- **ha_state** (String)


<a id="nestedblock--match--interface"></a>
### Nested Schema for `match.interface`

Optional:

- **ethernet** (Number)
- **loopback** (Number)
- **trunk** (Number)
- **tunnel** (Number)
- **ve** (Number)


<a id="nestedblock--match--ip"></a>
### Nested Schema for `match.ip`

Optional:

- **address** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ip--address))
- **next_hop** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ip--next_hop))
- **peer** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ip--peer))
- **rib** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ip--rib))

<a id="nestedblock--match--ip--address"></a>
### Nested Schema for `match.ip.address`

Optional:

- **acl1** (Number)
- **acl2** (Number)
- **name** (String)
- **prefix_list** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ip--address--prefix_list))

<a id="nestedblock--match--ip--address--prefix_list"></a>
### Nested Schema for `match.ip.address.prefix_list`

Optional:

- **name** (String)



<a id="nestedblock--match--ip--next_hop"></a>
### Nested Schema for `match.ip.next_hop`

Optional:

- **acl1** (Number)
- **acl2** (Number)
- **name** (String)
- **prefix_list_1** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ip--next_hop--prefix_list_1))

<a id="nestedblock--match--ip--next_hop--prefix_list_1"></a>
### Nested Schema for `match.ip.next_hop.prefix_list_1`

Optional:

- **name** (String)



<a id="nestedblock--match--ip--peer"></a>
### Nested Schema for `match.ip.peer`

Optional:

- **acl1** (Number)
- **acl2** (Number)
- **name** (String)


<a id="nestedblock--match--ip--rib"></a>
### Nested Schema for `match.ip.rib`

Optional:

- **exact** (String)
- **reachable** (String)
- **unreachable** (String)



<a id="nestedblock--match--ipv6"></a>
### Nested Schema for `match.ipv6`

Optional:

- **address_1** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ipv6--address_1))
- **next_hop_1** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ipv6--next_hop_1))
- **peer_1** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ipv6--peer_1))
- **rib** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ipv6--rib))

<a id="nestedblock--match--ipv6--address_1"></a>
### Nested Schema for `match.ipv6.address_1`

Optional:

- **name** (String)
- **prefix_list_2** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--ipv6--address_1--prefix_list_2))

<a id="nestedblock--match--ipv6--address_1--prefix_list_2"></a>
### Nested Schema for `match.ipv6.address_1.prefix_list_2`

Optional:

- **name** (String)



<a id="nestedblock--match--ipv6--next_hop_1"></a>
### Nested Schema for `match.ipv6.next_hop_1`

Optional:

- **next_hop_acl_name** (String)
- **prefix_list_name** (String)
- **v6_addr** (String)


<a id="nestedblock--match--ipv6--peer_1"></a>
### Nested Schema for `match.ipv6.peer_1`

Optional:

- **acl1** (Number)
- **acl2** (Number)
- **name** (String)


<a id="nestedblock--match--ipv6--rib"></a>
### Nested Schema for `match.ipv6.rib`

Optional:

- **exact** (String)
- **reachable** (String)
- **unreachable** (String)



<a id="nestedblock--match--local_preference"></a>
### Nested Schema for `match.local_preference`

Optional:

- **val** (Number)


<a id="nestedblock--match--metric"></a>
### Nested Schema for `match.metric`

Optional:

- **value** (Number)


<a id="nestedblock--match--origin"></a>
### Nested Schema for `match.origin`

Optional:

- **egp** (Number)
- **igp** (Number)
- **incomplete** (Number)


<a id="nestedblock--match--route_type"></a>
### Nested Schema for `match.route_type`

Optional:

- **external** (Block List, Max: 1) (see [below for nested schema](#nestedblock--match--route_type--external))

<a id="nestedblock--match--route_type--external"></a>
### Nested Schema for `match.route_type.external`

Optional:

- **value** (String)



<a id="nestedblock--match--scaleout"></a>
### Nested Schema for `match.scaleout`

Optional:

- **cluster_id** (Number)
- **operational_state** (String)


<a id="nestedblock--match--tag"></a>
### Nested Schema for `match.tag`

Optional:

- **value** (Number)



<a id="nestedblock--set"></a>
### Nested Schema for `set`

Optional:

- **aggregator** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--aggregator))
- **as_path** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--as_path))
- **atomic_aggregate** (Number)
- **comm_list** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--comm_list))
- **community** (String)
- **dampening_cfg** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--dampening_cfg))
- **ddos** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--ddos))
- **extcommunity** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--extcommunity))
- **ip** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--ip))
- **ipv6** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--ipv6))
- **level** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--level))
- **local_preference** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--local_preference))
- **metric** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--metric))
- **metric_type** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--metric_type))
- **origin** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--origin))
- **originator_id** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--originator_id))
- **tag** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--tag))
- **uuid** (String)
- **weight** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--weight))

<a id="nestedblock--set--aggregator"></a>
### Nested Schema for `set.aggregator`

Optional:

- **aggregator_as** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--aggregator--aggregator_as))

<a id="nestedblock--set--aggregator--aggregator_as"></a>
### Nested Schema for `set.aggregator.aggregator_as`

Optional:

- **asn** (Number)
- **ip** (String)



<a id="nestedblock--set--as_path"></a>
### Nested Schema for `set.as_path`

Optional:

- **num** (String)
- **num2** (String)
- **prepend** (String)


<a id="nestedblock--set--comm_list"></a>
### Nested Schema for `set.comm_list`

Optional:

- **delete** (Number)
- **name** (String)
- **name_delete** (Number)
- **v_exp** (Number)
- **v_exp_delete** (Number)
- **v_std** (Number)


<a id="nestedblock--set--dampening_cfg"></a>
### Nested Schema for `set.dampening_cfg`

Optional:

- **dampening** (Number)
- **dampening_half_time** (Number)
- **dampening_max_supress** (Number)
- **dampening_penalty** (Number)
- **dampening_reuse** (Number)
- **dampening_supress** (Number)


<a id="nestedblock--set--ddos"></a>
### Nested Schema for `set.ddos`

Optional:

- **class_list_cid** (Number)
- **class_list_name** (String)
- **zone** (String)


<a id="nestedblock--set--extcommunity"></a>
### Nested Schema for `set.extcommunity`

Optional:

- **rt** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--extcommunity--rt))
- **soo** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--extcommunity--soo))

<a id="nestedblock--set--extcommunity--rt"></a>
### Nested Schema for `set.extcommunity.rt`

Optional:

- **value** (String)


<a id="nestedblock--set--extcommunity--soo"></a>
### Nested Schema for `set.extcommunity.soo`

Optional:

- **value** (String)



<a id="nestedblock--set--ip"></a>
### Nested Schema for `set.ip`

Optional:

- **next_hop** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--ip--next_hop))

<a id="nestedblock--set--ip--next_hop"></a>
### Nested Schema for `set.ip.next_hop`

Optional:

- **address** (String)



<a id="nestedblock--set--ipv6"></a>
### Nested Schema for `set.ipv6`

Optional:

- **next_hop_1** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--ipv6--next_hop_1))

<a id="nestedblock--set--ipv6--next_hop_1"></a>
### Nested Schema for `set.ipv6.next_hop_1`

Optional:

- **address** (String)
- **local** (Block List, Max: 1) (see [below for nested schema](#nestedblock--set--ipv6--next_hop_1--local))

<a id="nestedblock--set--ipv6--next_hop_1--local"></a>
### Nested Schema for `set.ipv6.next_hop_1.local`

Optional:

- **address** (String)




<a id="nestedblock--set--level"></a>
### Nested Schema for `set.level`

Optional:

- **value** (String)


<a id="nestedblock--set--local_preference"></a>
### Nested Schema for `set.local_preference`

Optional:

- **val** (Number)


<a id="nestedblock--set--metric"></a>
### Nested Schema for `set.metric`

Optional:

- **value** (String)


<a id="nestedblock--set--metric_type"></a>
### Nested Schema for `set.metric_type`

Optional:

- **value** (String)


<a id="nestedblock--set--origin"></a>
### Nested Schema for `set.origin`

Optional:

- **egp** (Number)
- **igp** (Number)
- **incomplete** (Number)


<a id="nestedblock--set--originator_id"></a>
### Nested Schema for `set.originator_id`

Optional:

- **originator_ip** (String)


<a id="nestedblock--set--tag"></a>
### Nested Schema for `set.tag`

Optional:

- **value** (Number)


<a id="nestedblock--set--weight"></a>
### Nested Schema for `set.weight`

Optional:

- **weight_val** (Number)

