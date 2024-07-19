---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_ip_nat_translation Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_ip_nat_translation: Change or Disable NAT translation values
  PLACEHOLDER
---

# thunder_ip_nat_translation (Resource)

`thunder_ip_nat_translation`: Change or Disable NAT translation values

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_ip_nat_translation" "thunder_ip_nat_translation" {

  icmp_timeout {
    icmp_timeout = "fast"
  }
  ignore_tcp_msl = 1
  service_timeout_list {
    service_type = "tcp"
    port         = 57460
    timeout_type = "age"
    timeout_val  = 7756
  }
  tcp_timeout = 200
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `icmp_timeout` (Block List, Max: 1) (see [below for nested schema](#nestedblock--icmp_timeout))
- `ignore_tcp_msl` (Number) reclaim TCP resource immediately without MSL
- `service_timeout_list` (Block List) (see [below for nested schema](#nestedblock--service_timeout_list))
- `tcp_timeout` (Number) TCP protocol extended translations (Timeout in seconds (Interval of 60 seconds), default is 300 seconds (5 minutes))
- `udp_timeout` (Number) UDP protocol extended translations (Timeout in seconds (Interval of 60 seconds), default is 300 seconds (5 minutes))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--icmp_timeout"></a>
### Nested Schema for `icmp_timeout`

Optional:

- `icmp_timeout` (String) 'age': Expiration time; 'fast': Use Fast aging;
- `icmp_timeout_val` (Number) Timeout in seconds (Interval of 60 seconds)


<a id="nestedblock--service_timeout_list"></a>
### Nested Schema for `service_timeout_list`

Required:

- `port` (Number) Port Number
- `service_type` (String) 'tcp': TCP Protocol; 'udp': UDP Protocol;

Optional:

- `timeout_type` (String) 'age': Expiration time; 'fast': Use Fast aging;
- `timeout_val` (Number) Timeout in seconds (Interval of 60 seconds)
- `uuid` (String) uuid of the object

