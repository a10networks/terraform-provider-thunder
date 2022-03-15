---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_snmp_server_disable_traps Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_snmp_server_disable_traps: Disable l3v partition SNMP traps
  This resource is private-partition-only
---

# thunder_snmp_server_disable_traps (Resource)

`thunder_snmp_server_disable_traps`: Disable l3v partition SNMP traps

This resource is private-partition-only



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **all** (Number) Disable all traps on this partition
- **gslb** (Number) Disable all gslb traps on this partition
- **id** (String) The ID of this resource.
- **slb** (Number) Disable all slb traps on this partition
- **slb_change** (Number) Disable all slb-change traps on this partition
- **snmp** (Number) Disable all snmp traps on this partition
- **uuid** (String) uuid of the object
- **vrrp_a** (Number) Disable all vrrp-a on this partition

