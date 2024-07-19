---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_cgnv6_lsn_alg_ftp_stats Data Source - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_cgnv6_lsn_alg_ftp_stats: Statistics for the object ftp
  PLACEHOLDER
---

# thunder_cgnv6_lsn_alg_ftp_stats (Data Source)

`thunder_cgnv6_lsn_alg_ftp_stats`: Statistics for the object ftp

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
data "thunder_cgnv6_lsn_alg_ftp_stats" "thunder_cgnv6_lsn_alg_ftp_stats" {
}
output "get_cgnv6_lsn_alg_ftp_stats" {
  value = ["${data.thunder_cgnv6_lsn_alg_ftp_stats.thunder_cgnv6_lsn_alg_ftp_stats}"]
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

- `eprt_requests` (Number) EPRT Requests From Client
- `epsv_replies` (Number) EPSV Replies From Server
- `lprt_requests` (Number) LPRT Requests From Client
- `lpsv_replies` (Number) LPSV Replies From Server
- `pasv_replies` (Number) PASV Replies From Server
- `port_requests` (Number) PORT Requests From Client

