---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_scaleout_cluster_local_device Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_scaleout_cluster_local_device: Local device configuration
  PLACEHOLDER
---

# thunder_scaleout_cluster_local_device (Resource)

`thunder_scaleout_cluster_local_device`: Local device configuration

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_scaleout_cluster_local_device" "thunder_scaleout_cluster_local_device" {

  cluster_id  = 2
  action      = "enable"
  priority    = 209
  start_delay = 167
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_id` (String) ClusterId

### Optional

- `action` (String) 'enable': enable; 'disable': disable;
- `cluster_mode` (String) 'layer-2': Nodes in cluster are layer 2 connected (default mode); 'layer-3': Nodes in cluster are l3 connected;
- `exclude_interfaces` (Block List, Max: 1) (see [below for nested schema](#nestedblock--exclude_interfaces))
- `id1` (Number)
- `l2_redirect` (Block List, Max: 1) (see [below for nested schema](#nestedblock--l2_redirect))
- `priority` (Number)
- `session_sync` (Block List, Max: 1) (see [below for nested schema](#nestedblock--session_sync))
- `start_delay` (Number)
- `tracking_template` (Block List, Max: 1) (see [below for nested schema](#nestedblock--tracking_template))
- `traffic_redirection` (Block List, Max: 1) (see [below for nested schema](#nestedblock--traffic_redirection))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--exclude_interfaces"></a>
### Nested Schema for `exclude_interfaces`

Optional:

- `eth_cfg` (Block List) (see [below for nested schema](#nestedblock--exclude_interfaces--eth_cfg))
- `loopback_cfg` (Block List) (see [below for nested schema](#nestedblock--exclude_interfaces--loopback_cfg))
- `trunk_cfg` (Block List) (see [below for nested schema](#nestedblock--exclude_interfaces--trunk_cfg))
- `uuid` (String) uuid of the object
- `ve_cfg` (Block List) (see [below for nested schema](#nestedblock--exclude_interfaces--ve_cfg))

<a id="nestedblock--exclude_interfaces--eth_cfg"></a>
### Nested Schema for `exclude_interfaces.eth_cfg`

Optional:

- `ethernet` (Number) Ethernet Interface (Ethernet interface number)


<a id="nestedblock--exclude_interfaces--loopback_cfg"></a>
### Nested Schema for `exclude_interfaces.loopback_cfg`

Optional:

- `loopback` (Number) Loopback Interface (Loopback interface number)


<a id="nestedblock--exclude_interfaces--trunk_cfg"></a>
### Nested Schema for `exclude_interfaces.trunk_cfg`

Optional:

- `trunk` (Number) Trunk Interface (Trunk interface number)


<a id="nestedblock--exclude_interfaces--ve_cfg"></a>
### Nested Schema for `exclude_interfaces.ve_cfg`

Optional:

- `ve` (Number) Virtual ethernet Interface (Virtual ethernet interface number)



<a id="nestedblock--l2_redirect"></a>
### Nested Schema for `l2_redirect`

Optional:

- `ethernet_vlan` (Number) VLAN ID
- `redirect_eth` (Number) Ethernet port (Port Value)
- `redirect_trunk` (Number) L2 Trunk group
- `trunk_vlan` (Number) VLAN ID
- `uuid` (String) uuid of the object


<a id="nestedblock--session_sync"></a>
### Nested Schema for `session_sync`

Optional:

- `follow_shared` (Number) Follow shared partition for session sync
- `interfaces` (Block List, Max: 1) (see [below for nested schema](#nestedblock--session_sync--interfaces))
- `reachability_options` (Block List, Max: 1) (see [below for nested schema](#nestedblock--session_sync--reachability_options))
- `uuid` (String) uuid of the object

<a id="nestedblock--session_sync--interfaces"></a>
### Nested Schema for `session_sync.interfaces`

Optional:

- `eth_cfg` (Block List) (see [below for nested schema](#nestedblock--session_sync--interfaces--eth_cfg))
- `loopback_cfg` (Block List) (see [below for nested schema](#nestedblock--session_sync--interfaces--loopback_cfg))
- `trunk_cfg` (Block List) (see [below for nested schema](#nestedblock--session_sync--interfaces--trunk_cfg))
- `uuid` (String) uuid of the object
- `ve_cfg` (Block List) (see [below for nested schema](#nestedblock--session_sync--interfaces--ve_cfg))

<a id="nestedblock--session_sync--interfaces--eth_cfg"></a>
### Nested Schema for `session_sync.interfaces.eth_cfg`

Optional:

- `ethernet` (Number) Ethernet Interface (Ethernet interface number)


<a id="nestedblock--session_sync--interfaces--loopback_cfg"></a>
### Nested Schema for `session_sync.interfaces.loopback_cfg`

Optional:

- `loopback` (Number) Loopback Interface(Not applicable in 'layer-2' mode) (Loopback interface number)


<a id="nestedblock--session_sync--interfaces--trunk_cfg"></a>
### Nested Schema for `session_sync.interfaces.trunk_cfg`

Optional:

- `trunk` (Number) Trunk Interface (Trunk interface number)


<a id="nestedblock--session_sync--interfaces--ve_cfg"></a>
### Nested Schema for `session_sync.interfaces.ve_cfg`

Optional:

- `ve` (Number) Virtual ethernet Interface (Virtual ethernet interface number)



<a id="nestedblock--session_sync--reachability_options"></a>
### Nested Schema for `session_sync.reachability_options`

Optional:

- `skip_default_route` (Number) Do not choose default route for redirection(Not applicable in 'layer-2' mode)
- `uuid` (String) uuid of the object



<a id="nestedblock--tracking_template"></a>
### Nested Schema for `tracking_template`

Optional:

- `multi_template_list` (Block List) (see [below for nested schema](#nestedblock--tracking_template--multi_template_list))
- `template_list` (Block List) (see [below for nested schema](#nestedblock--tracking_template--template_list))

<a id="nestedblock--tracking_template--multi_template_list"></a>
### Nested Schema for `tracking_template.multi_template_list`

Required:

- `multi_template` (String) bind multi tracking template name

Optional:

- `action` (String) 'down': node stops processing user traffic; 'exit-cluster': node exits scaleout cluster;
- `template` (Block List) (see [below for nested schema](#nestedblock--tracking_template--multi_template_list--template))
- `threshold` (Number) action triggering threshold
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--tracking_template--multi_template_list--template"></a>
### Nested Schema for `tracking_template.multi_template_list.template`

Optional:

- `partition_name` (String) Partition name
- `template_name` (String) bind tracking template name



<a id="nestedblock--tracking_template--template_list"></a>
### Nested Schema for `tracking_template.template_list`

Required:

- `template` (String) bind tracking template name

Optional:

- `threshold_cfg` (Block List) (see [below for nested schema](#nestedblock--tracking_template--template_list--threshold_cfg))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--tracking_template--template_list--threshold_cfg"></a>
### Nested Schema for `tracking_template.template_list.threshold_cfg`

Optional:

- `action` (String) 'down': node stops processing user traffic; 'exit-cluster': node exits scaleout cluster;
- `threshold` (Number) action triggering threshold




<a id="nestedblock--traffic_redirection"></a>
### Nested Schema for `traffic_redirection`

Optional:

- `follow_shared` (Number) Follow shared partition for redirection
- `interfaces` (Block List, Max: 1) (see [below for nested schema](#nestedblock--traffic_redirection--interfaces))
- `reachability_options` (Block List, Max: 1) (see [below for nested schema](#nestedblock--traffic_redirection--reachability_options))
- `uuid` (String) uuid of the object

<a id="nestedblock--traffic_redirection--interfaces"></a>
### Nested Schema for `traffic_redirection.interfaces`

Optional:

- `eth_cfg` (Block List) (see [below for nested schema](#nestedblock--traffic_redirection--interfaces--eth_cfg))
- `loopback_cfg` (Block List) (see [below for nested schema](#nestedblock--traffic_redirection--interfaces--loopback_cfg))
- `trunk_cfg` (Block List) (see [below for nested schema](#nestedblock--traffic_redirection--interfaces--trunk_cfg))
- `uuid` (String) uuid of the object
- `ve_cfg` (Block List) (see [below for nested schema](#nestedblock--traffic_redirection--interfaces--ve_cfg))

<a id="nestedblock--traffic_redirection--interfaces--eth_cfg"></a>
### Nested Schema for `traffic_redirection.interfaces.eth_cfg`

Optional:

- `ethernet` (Number) Ethernet Interface (Ethernet interface number)


<a id="nestedblock--traffic_redirection--interfaces--loopback_cfg"></a>
### Nested Schema for `traffic_redirection.interfaces.loopback_cfg`

Optional:

- `loopback` (Number) Loopback Interface (Loopback interface number)


<a id="nestedblock--traffic_redirection--interfaces--trunk_cfg"></a>
### Nested Schema for `traffic_redirection.interfaces.trunk_cfg`

Optional:

- `trunk` (Number) Trunk Interface (Trunk interface number)


<a id="nestedblock--traffic_redirection--interfaces--ve_cfg"></a>
### Nested Schema for `traffic_redirection.interfaces.ve_cfg`

Optional:

- `ve` (Number) Virtual ethernet Interface (Virtual ethernet interface number)



<a id="nestedblock--traffic_redirection--reachability_options"></a>
### Nested Schema for `traffic_redirection.reachability_options`

Optional:

- `skip_default_route` (Number) Do not choose default route for redirection
- `uuid` (String) uuid of the object

