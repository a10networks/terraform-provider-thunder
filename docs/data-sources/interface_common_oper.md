---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_interface_common_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_interface_common_oper: Operational Status for the object common
  PLACEHOLDER
---

# thunder_interface_common_oper (Data Source)

`thunder_interface_common_oper`: Operational Status for the object common

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_interface_common_oper" "thunder_interface_common_oper" {
}
output "get_interface_common_oper" {
  value = ["${data.thunder_interface_common_oper.thunder_interface_common_oper}"]
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

- `interfaces` (Block List) (see [below for nested schema](#nestedblock--oper--interfaces))
- `interval` (Number)
- `time` (String)
- `tot_num_phy_intf` (Number)
- `vnp_id` (Number)

<a id="nestedblock--oper--interfaces"></a>
### Nested Schema for `oper.interfaces`

Optional:

- `input_utilization` (Number)
- `is_part_default_vlan` (String)
- `output_utilization` (Number)
- `port_num` (String)
- `port_type` (String)
- `rate_byte_rcvd` (Number)
- `rate_byte_sent` (Number)
- `rate_pkt_rcvd` (Number)
- `rate_pkt_sent` (Number)
- `span_tree` (String)
- `total_lane` (Number)
- `transceivers_info` (Block List) (see [below for nested schema](#nestedblock--oper--interfaces--transceivers_info))
- `type_vendor_part` (String)

<a id="nestedblock--oper--interfaces--transceivers_info"></a>
### Nested Schema for `oper.interfaces.transceivers_info`

Optional:

- `curr` (String)
- `hi_alarm` (String)
- `hi_warn` (String)
- `lane` (Number)
- `lo_alarm` (String)
- `lo_warn` (String)
- `transceiver_type` (String)

