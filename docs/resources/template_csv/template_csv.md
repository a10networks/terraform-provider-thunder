---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_template_csv Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_template_csv: Specify csv template
  PLACEHOLDER
---

# thunder_template_csv (Resource)

`thunder_template_csv`: Specify csv template

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_template_csv" "thunder_template_csv" {
  csv_name    = "test"
  delim_char  = ","
  ipv6_enable = 1
  multiple_fields {
    field    = 17
    csv_type = "ip-from"
  }
  user_tag = "62"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `csv_name` (String) Specify name of csv template

### Optional

- `delim_char` (String) enter a delimiter character, default ","
- `delim_num` (Number) enter a delimiter number, default 44 (",")
- `ipv6_enable` (Number) Support IPv6 IP ranges
- `multiple_fields` (Block List) (see [below for nested schema](#nestedblock--multiple_fields))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--multiple_fields"></a>
### Nested Schema for `multiple_fields`

Optional:

- `csv_type` (String) 'ip-from': Beginning address of IP range or subnet; 'ip-to-mask': Ending address of IP range or Mask; 'continent': Continent; 'country': Country; 'state': State or province; 'city': City;
- `field` (Number) Field index number (Index of Field)

