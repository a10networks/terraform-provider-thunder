---
layout: "thunder"
page_title: "thunder: thunder_slb_template_csv"
sidebar_current: "docs-thunder-resource-slb-template-csv"
description: |-
    Provides details about thunder slb template csv resource for A10
---

# thunder\_slb\_template\_csv

`thunder_slb_template_csv` provides details about slb template csv
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_csv" "csv" {
	csv_name = "testcsv"
	user_tag = "test_tag"
	ipv6_enable = 0
	delim_num = 0
	multiple_fields {
		field = 5 
		csv_type = "ip-from"
	}
}
```

## Argument Reference

* `csv_name` - Specify name of csv template
* `user_tag` - Customized tag
* `ipv6_enable` - Support IPv6 IP ranges
* `delim_num` - enter a delimiter number, default 44 (“,”)
* `multiple_fields` -
    * `field` - Field index number (Index of Field)
    * `csv_type` - 'ip-from’: Beginning address of IP range or subnet; 'ip-to-mask’: Ending address of IP range or Mask; 'continent’: Continent; 'country’: Country; 'state’: State or province; 'city’: City;

