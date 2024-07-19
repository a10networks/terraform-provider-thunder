---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_fix_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_fix_stats: Statistics for the object fix
  PLACEHOLDER
---

# thunder_slb_fix_stats (Data Source)

`thunder_slb_fix_stats`: Statistics for the object fix

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_slb_fix_stats" "thunder_slb_fix_stats" {
}
output "get_slb_fix_stats" {
  value = ["${data.thunder_slb_fix_stats.thunder_slb_fix_stats}"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `stats` (Block List, Max: 1) (see [below for nested schema](#nestedblock--stats))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--stats"></a>
### Nested Schema for `stats`

Optional:

- `client_err` (Number) Client fail
- `client_tls_conn` (Number) Client TLS conn
- `curr_proxy` (Number) Current proxy conns
- `default_switching` (Number) Default switching
- `insert_clientip` (Number) Insert client IP
- `noroute` (Number) No route failure
- `sender_switching` (Number) Sender ID switching
- `server_err` (Number) Server fail
- `server_tls_conn` (Number) Server TLS conn
- `snat_fail` (Number) Source NAT failure
- `svrsel_fail` (Number) Server selection failure
- `target_switching` (Number) Target ID switching
- `total_proxy` (Number) Total proxy conns

