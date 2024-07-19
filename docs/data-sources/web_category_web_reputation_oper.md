---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_web_category_web_reputation_oper Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_web_category_web_reputation_oper: Operational Status for the object web-reputation
  PLACEHOLDER
---

# thunder_web_category_web_reputation_oper (Data Source)

`thunder_web_category_web_reputation_oper`: Operational Status for the object web-reputation

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_web_category_web_reputation_oper" "thunder_web_category_web_reputation_oper" {

}
output "get_web_category_web_reputation_oper" {
  value = ["${data.thunder_web_category_web_reputation_oper.thunder_web_category_web_reputation_oper}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `bypassed_urls` (Block List, Max: 1) (see [below for nested schema](#nestedblock--bypassed_urls))
- `intercepted_urls` (Block List, Max: 1) (see [below for nested schema](#nestedblock--intercepted_urls))
- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--oper))
- `url` (Block List, Max: 1) (see [below for nested schema](#nestedblock--url))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--bypassed_urls"></a>
### Nested Schema for `bypassed_urls`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--bypassed_urls--oper))

<a id="nestedblock--bypassed_urls--oper"></a>
### Nested Schema for `bypassed_urls.oper`

Optional:

- `all_urls` (String)
- `number_of_urls` (Number)
- `url_list` (Block List) (see [below for nested schema](#nestedblock--bypassed_urls--oper--url_list))
- `url_name` (String)

<a id="nestedblock--bypassed_urls--oper--url_list"></a>
### Nested Schema for `bypassed_urls.oper.url_list`

Optional:

- `url_name` (String)




<a id="nestedblock--intercepted_urls"></a>
### Nested Schema for `intercepted_urls`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--intercepted_urls--oper))

<a id="nestedblock--intercepted_urls--oper"></a>
### Nested Schema for `intercepted_urls.oper`

Optional:

- `all_urls` (String)
- `number_of_urls` (Number)
- `url_list` (Block List) (see [below for nested schema](#nestedblock--intercepted_urls--oper--url_list))
- `url_name` (String)

<a id="nestedblock--intercepted_urls--oper--url_list"></a>
### Nested Schema for `intercepted_urls.oper.url_list`

Optional:

- `url_name` (String)




<a id="nestedblock--oper"></a>
### Nested Schema for `oper`

Optional:

- `all_urls` (String)
- `number_of_urls` (Number)
- `url_list` (Block List) (see [below for nested schema](#nestedblock--oper--url_list))
- `url_name` (String)

<a id="nestedblock--oper--url_list"></a>
### Nested Schema for `oper.url_list`

Optional:

- `url_name` (String)



<a id="nestedblock--url"></a>
### Nested Schema for `url`

Optional:

- `oper` (Block List, Max: 1) (see [below for nested schema](#nestedblock--url--oper))

<a id="nestedblock--url--oper"></a>
### Nested Schema for `url.oper`

Optional:

- `local_db_only` (Number)
- `name` (String)
- `reputation_score` (String)

