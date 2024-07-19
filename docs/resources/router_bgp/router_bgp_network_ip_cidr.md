---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_router_bgp_network_ip_cidr Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_router_bgp_network_ip_cidr: Specify a ip network to announce via BGP
  PLACEHOLDER
---

# thunder_router_bgp_network_ip_cidr (Resource)

`thunder_router_bgp_network_ip_cidr`: Specify a ip network to announce via BGP

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_router_bgp_network_ip_cidr" "thunder_router_bgp_network_ip_cidr" {

  as_number         = "300"
  network_ipv4_cidr = "19.0.2.0/24"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `as_number` (String) AsNumber
- `network_ipv4_cidr` (String) Specify network mask

### Optional

- `backdoor` (Number) Specify a BGP backdoor route
- `comm_value` (String) community value in the format 1-4294967295|AA:NN|internet|local-AS|no-advertise|no-export
- `description` (String) Network specific description (Up to 80 characters describing this network)
- `lcomm_value` (String) Large community value in the format XX:YY:ZZ
- `route_map` (String) Route-map to modify the attributes (Name of the route map)
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

