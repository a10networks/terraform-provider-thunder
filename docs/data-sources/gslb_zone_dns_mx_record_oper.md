---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_zone_dns_mx_record_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_zone_dns_mx_record_oper: Operational Status for the object dns-mx-record
  PLACEHOLDER
---

# thunder_gslb_zone_dns_mx_record_oper (Data Source)

`thunder_gslb_zone_dns_mx_record_oper`: Operational Status for the object dns-mx-record

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_zone_dns_mx_record_oper" "thunder_gslb_zone_dns_mx_record_oper" {

  name    = "a11"
  mx_name = "69"
}
output "get_gslb_zone_dns_mx_record_oper" {
  value = ["${data.thunder_gslb_zone_dns_mx_record_oper.thunder_gslb_zone_dns_mx_record_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `mx_name` (String) Specify Domain Name
- `name` (String) Name

### Optional

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `hits` (Number)
- `last_server` (String)
- `priority` (Number)

