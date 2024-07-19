---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_zone_service_dns_record Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_zone_service_dns_record: Specify DNS Record
  PLACEHOLDER
---

# thunder_gslb_zone_service_dns_record (Resource)

`thunder_gslb_zone_service_dns_record`: Specify DNS Record

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service_dns_record" "thunder_gslb_zone_service_dns_record" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  data         = "442"
  type         = 26622
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name
- `service_name` (String) ServiceName
- `service_port` (String) ServicePort
- `type` (Number) Specify DNS Type

### Optional

- `data` (String) Specify DNS Data
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

