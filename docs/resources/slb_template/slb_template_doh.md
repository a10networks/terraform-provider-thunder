---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_template_doh Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_template_doh: DNS over HTTP(s) template
  PLACEHOLDER
---

# thunder_slb_template_doh (Resource)

`thunder_slb_template_doh`: DNS over HTTP(s) template

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_doh" "thunder_slb_template_doh" {
  name       = "test1"
  conn_reuse = "disable"
  dns        = "default"
  dns_retry {
    retry_interval = 10
    after_timeout  = "close"
    max_trials     = 3
  }
  forwarder {
    v4_internal = 0
    v4_port     = 53
    v4_l4_proto = "both"
    v6_internal = 0
    v6_port     = 53
    v6_l4_proto = "both"
    bypass_doh  = 0
  }
  non_dns_request                     = "reject"
  reject_status_code                  = "400"
  shared_partition_tcp_proxy_template = 0
  user_tag                            = "77"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) DNS over HTTP(s) Template Name

### Optional

- `conn_reuse` (String) 'enable': Enable Connection Reuse; 'disable': Disable Connection-Reuse (Default);
- `dns` (String) DNS Template Name
- `dns_retry` (Block List, Max: 1) (see [below for nested schema](#nestedblock--dns_retry))
- `forwarder` (Block List, Max: 1) (see [below for nested schema](#nestedblock--forwarder))
- `non_dns_request` (String) 'allow': Forward Non-DoH request to http server bound to vport; 'reject': Reject Non-DoH requests with HTTP 400 Bad Request (Default);
- `reject_status_code` (String) '400': Status Code 400 BAD Request (Default); '500': Status Code 500 Internal Server Error; '501': Status Code 501 Not Implemented;
- `shared_partition_dns_template` (Number) Reference a DNS template from shared partition
- `shared_partition_tcp_proxy_template` (Number) Reference a TCP Proxy template from shared partition
- `snat_pool` (String) Source NAT pool or pool group
- `source_nat` (String) 'auto': Perform Source NAT Auto for service-group(Default) (Not supported with forwarding-ip); 'disable': Don't perform source-nat for server side DNS queries; 'pool': Perform Source NAT with specific pool;
- `tcp_proxy` (String) TCP Proxy Template Name
- `template_dns_shared` (String) DNS Template name
- `template_tcp_proxy_shared` (String) TCP Proxy Template name
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--dns_retry"></a>
### Nested Schema for `dns_retry`

Optional:

- `after_timeout` (String) 'close': Close client side connection; 'retry-with-tcp': Retry DNS query to server using TCP (If UDP was tried initially. Close after.);
- `max_trials` (Number) Total number of times to try DNS query to server before closing client connection, default 3
- `retry_interval` (Number) DNS Retry Interval value 1 - 400 in units of 100ms, default is 10 (default is 1000ms) (1 - 400 in units of 100ms, default is 10 (1000ms/1sec))
- `uuid` (String) uuid of the object


<a id="nestedblock--forwarder"></a>
### Nested Schema for `forwarder`

Optional:

- `bypass_doh` (Number) Forward valid DoH HTTP request as is, no DNS packet extraction (Bypass DoH)
- `forwarding_ipv4` (String) SLB VIP IPv4 address to forward DOH query (IP address)
- `forwarding_ipv6` (String) SLB VIP IPv6 address to forward DOH query (IP address)
- `tcp_service_group` (String) Bind a TCP Service Group to the template (Service Group Name)
- `udp_service_group` (String) Bind a UDP Service Group to the template (Service Group Name)
- `uuid` (String) uuid of the object
- `v4_internal` (Number) Try to find this IP as a VIP in this L3v Partition and forward it internally to the VIP
- `v4_l4_proto` (String) 'tcp': Use TCP only when forwarding DNS traffic; 'udp': Use UDP only when forwarding DNS traffic; 'both': Use UDP 1st and if unreachable, retry with TCP when forwarding DNS traffic;
- `v4_port` (Number) Forwarding port number, Default is 53
- `v6_internal` (Number) Try to find this IP as a VIP in this L3v Partition and forward it internally to the VIP
- `v6_l4_proto` (String) 'tcp': Use TCP only when forwarding DNS traffic; 'udp': Use UDP only when forwarding DNS traffic; 'both': Use UDP 1st and if unreachable, retry with TCP when forwarding DNS traffic;
- `v6_port` (Number) Forwarding port number, Default is 53

