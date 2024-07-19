---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_site_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_site_stats: Statistics for the object site
  PLACEHOLDER
---

# thunder_gslb_site_stats (Data Source)

`thunder_gslb_site_stats`: Statistics for the object site

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_gslb_site_stats" "thunder_gslb_site_stats" {

  site_name = "3"
}
output "get_gslb_site_stats" {
  value = ["${data.thunder_gslb_site_stats.thunder_gslb_site_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `site_name` (String) Specify GSLB site name

### Optional

- `ip_server_list` (Block List) (see [below for nested schema](#nestedblock--ip_server_list))
- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--ip_server_list"></a>
### Nested Schema for `ip_server_list`

Required:

- `ip_server_name` (String) Specify the real server name

Optional:

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--ip_server_list--stats))

<a id="nestedblock--ip_server_list--stats"></a>
### Nested Schema for `ip_server_list.stats`

Optional:

- `hits` (Number) Number of times the IP was selected
- `recent` (Number) Recent hits



<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `hits` (Number) Number of times the site was selected

