---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_gslb_zone_service Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_gslb_zone_service: Service information for the GSLB zone
  PLACEHOLDER
---

# thunder_gslb_zone_service (Resource)

`thunder_gslb_zone_service`: Service information for the GSLB zone

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_gslb_zone_service" "thunder_gslb_zone_service" {

  name         = "a11"
  service_name = "s25"
  service_port = 33159
  action       = "forward"
  forward_type = "both"
  disable      = 0
  dns_a_record {
    dns_a_record_srv_list {
      svrname    = "test-server1"
      no_resp    = 0
      as_backup  = 0
      weight     = 96
      ttl        = 0
      as_replace = 0
      disable    = 1
      static     = 1
      admin_ip   = 198
    }
    dns_a_record_ipv4_list {
      dns_a_record_ip = "10.10.10.10"
      no_resp         = 0
      as_backup       = 0
      weight          = 31
      ttl             = 0
      as_replace      = 0
      disable         = 0
      static          = 1
      admin_ip        = 232
    }
    dns_a_record_ipv6_list {
      dns_a_record_ipv6 = "2001:db8:3333:4444:5555:6666:7777:8888"
      no_resp           = 0
      as_backup         = 0
      weight            = 97
      ttl               = 0
      as_replace        = 0
      disable           = 0
      static            = 1
      admin_ip          = 202
    }
  }
  dns_caa_record_list {
    critical_flag = 106
    property_tag  = "32"
    rdata         = "798"
    ttl           = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_cname_record_list {
    alias_name       = "89"
    admin_preference = 100
    weight           = 1
    as_backup        = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_mx_record_list {
    mx_name  = "112"
    priority = 35640
    ttl      = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_naptr_record_list {
    naptr_target  = "35"
    service_proto = "68"
    flag          = "1"
    order         = 8489
    preference    = 16018
    regexp        = 0
    ttl           = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_ns_record_list {
    ns_name = "12"
    ttl     = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_ptr_record_list {
    ptr_name = "125"
    ttl      = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_record_list {
    type = 57856
    data = "438"
  }
  dns_srv_record_list {
    srv_name = "70"
    port     = 27600
    priority = 58080
    weight   = 10
    ttl      = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  dns_txt_record_list {
    record_name = "11"
    txt_data    = "126"
    ttl         = 0
    sampling_enable {
      counters1 = "all"
    }
  }
  geo_location_list {
    geo_name = "75"
    user_tag = "70"
  }
  health_check_gateway = "enable"
  health_check_port {
    health_check_port = 37739
  }
  sampling_enable {
    counters1 = "all"
  }
  user_tag = "56"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name
- `service_name` (String) Specify the service name for the zone, * for wildcard
- `service_port` (Number) Port number of the service

### Optional

- `action` (String) 'drop': Drop query; 'forward': Forward packet; 'ignore': Send empty response; 'reject': Send refuse response;
- `disable` (Number) Disable
- `dns_a_record` (Block List, Max: 1) (see [below for nested schema](#nestedblock--dns_a_record))
- `dns_caa_record_list` (Block List) (see [below for nested schema](#nestedblock--dns_caa_record_list))
- `dns_cname_record_list` (Block List) (see [below for nested schema](#nestedblock--dns_cname_record_list))
- `dns_mx_record_list` (Block List) (see [below for nested schema](#nestedblock--dns_mx_record_list))
- `dns_naptr_record_list` (Block List) (see [below for nested schema](#nestedblock--dns_naptr_record_list))
- `dns_ns_record_list` (Block List) (see [below for nested schema](#nestedblock--dns_ns_record_list))
- `dns_ptr_record_list` (Block List) (see [below for nested schema](#nestedblock--dns_ptr_record_list))
- `dns_record_list` (Block List) (see [below for nested schema](#nestedblock--dns_record_list))
- `dns_srv_record_list` (Block List) (see [below for nested schema](#nestedblock--dns_srv_record_list))
- `dns_txt_record_list` (Block List) (see [below for nested schema](#nestedblock--dns_txt_record_list))
- `forward_type` (String) 'both': Forward both query and response; 'query': Forward query; 'response': Forward response;
- `geo_location_list` (Block List) (see [below for nested schema](#nestedblock--geo_location_list))
- `health_check_gateway` (String) 'enable': Enable Gateway Status Check; 'disable': Disable Gateway Status Check;
- `health_check_port` (Block List) (see [below for nested schema](#nestedblock--health_check_port))
- `policy` (String) Specify policy for this service (Specify policy name)
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--sampling_enable))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--dns_a_record"></a>
### Nested Schema for `dns_a_record`

Optional:

- `dns_a_record_ipv4_list` (Block List) (see [below for nested schema](#nestedblock--dns_a_record--dns_a_record_ipv4_list))
- `dns_a_record_ipv6_list` (Block List) (see [below for nested schema](#nestedblock--dns_a_record--dns_a_record_ipv6_list))
- `dns_a_record_srv_list` (Block List) (see [below for nested schema](#nestedblock--dns_a_record--dns_a_record_srv_list))

<a id="nestedblock--dns_a_record--dns_a_record_ipv4_list"></a>
### Nested Schema for `dns_a_record.dns_a_record_ipv4_list`

Required:

- `dns_a_record_ip` (String) Specify IP address

Optional:

- `admin_ip` (Number) Specify admin priority of Service-IP (Specify the priority)
- `as_backup` (Number) As backup when fail
- `as_replace` (Number) Return this Service-IP when enable ip-replace
- `disable` (Number) Disable this Service-IP
- `no_resp` (Number) Don't use this Service-IP as DNS response
- `static` (Number) Return this Service-IP in DNS server mode
- `ttl` (Number) Specify TTL for Service-IP
- `uuid` (String) uuid of the object
- `weight` (Number) Specify weight for Service-IP (Weight value)


<a id="nestedblock--dns_a_record--dns_a_record_ipv6_list"></a>
### Nested Schema for `dns_a_record.dns_a_record_ipv6_list`

Required:

- `dns_a_record_ipv6` (String) IPV6 address

Optional:

- `admin_ip` (Number) Specify admin priority of Service-IP (Specify the priority)
- `as_backup` (Number) As backup when fail
- `as_replace` (Number) Return this Service-IP when enable ip-replace
- `disable` (Number) Disable this Service-IP
- `no_resp` (Number) Don't use this Service-IP as DNS response
- `static` (Number) Return this Service-IP in DNS server mode
- `ttl` (Number) Specify TTL for Service-IP
- `uuid` (String) uuid of the object
- `weight` (Number) Specify weight for Service-IP (Weight value)


<a id="nestedblock--dns_a_record--dns_a_record_srv_list"></a>
### Nested Schema for `dns_a_record.dns_a_record_srv_list`

Required:

- `svrname` (String) Specify name

Optional:

- `admin_ip` (Number) Specify admin priority of Service-IP (Specify the priority)
- `as_backup` (Number) As backup when fail
- `as_replace` (Number) Return this Service-IP when enable ip-replace
- `disable` (Number) Disable this Service-IP
- `no_resp` (Number) Don't use this Service-IP as DNS response
- `static` (Number) Return this Service-IP in DNS server mode
- `ttl` (Number) Specify TTL for Service-IP
- `uuid` (String) uuid of the object
- `weight` (Number) Specify weight for Service-IP (Weight value)



<a id="nestedblock--dns_caa_record_list"></a>
### Nested Schema for `dns_caa_record_list`

Required:

- `critical_flag` (Number) Issuer Critical Flag
- `property_tag` (String) Specify other property tags, only allowed lowercase alphanumeric
- `rdata` (String) Specify the Issuer Domain Name or a URL

Optional:

- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--dns_caa_record_list--sampling_enable))
- `ttl` (Number) Specify TTL
- `uuid` (String) uuid of the object

<a id="nestedblock--dns_caa_record_list--sampling_enable"></a>
### Nested Schema for `dns_caa_record_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'hits': Number of times the CAA has been used;



<a id="nestedblock--dns_cname_record_list"></a>
### Nested Schema for `dns_cname_record_list`

Required:

- `alias_name` (String) Specify the alias name

Optional:

- `admin_preference` (Number) Specify Administrative Preference, default is 100
- `as_backup` (Number) As backup when fail
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--dns_cname_record_list--sampling_enable))
- `uuid` (String) uuid of the object
- `weight` (Number) Specify Weight, default is 1

<a id="nestedblock--dns_cname_record_list--sampling_enable"></a>
### Nested Schema for `dns_cname_record_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'cname-hits': Number of times the CNAME has been used;



<a id="nestedblock--dns_mx_record_list"></a>
### Nested Schema for `dns_mx_record_list`

Required:

- `mx_name` (String) Specify Domain Name

Optional:

- `priority` (Number) Specify Priority
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--dns_mx_record_list--sampling_enable))
- `ttl` (Number) Specify TTL
- `uuid` (String) uuid of the object

<a id="nestedblock--dns_mx_record_list--sampling_enable"></a>
### Nested Schema for `dns_mx_record_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'hits': Number of times the record has been used;



<a id="nestedblock--dns_naptr_record_list"></a>
### Nested Schema for `dns_naptr_record_list`

Required:

- `flag` (String) Specify the flag (e.g., a, s). Default is empty flag
- `naptr_target` (String) Specify the replacement or regular expression
- `service_proto` (String) Specify Service and Protocol

Optional:

- `order` (Number) Specify Order
- `preference` (Number) Specify Preference
- `regexp` (Number) Return the regular expression
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--dns_naptr_record_list--sampling_enable))
- `ttl` (Number) Specify TTL
- `uuid` (String) uuid of the object

<a id="nestedblock--dns_naptr_record_list--sampling_enable"></a>
### Nested Schema for `dns_naptr_record_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'naptr-hits': Number of times the NAPTR has been used;



<a id="nestedblock--dns_ns_record_list"></a>
### Nested Schema for `dns_ns_record_list`

Required:

- `ns_name` (String) Specify Domain Name

Optional:

- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--dns_ns_record_list--sampling_enable))
- `ttl` (Number) Specify TTL
- `uuid` (String) uuid of the object

<a id="nestedblock--dns_ns_record_list--sampling_enable"></a>
### Nested Schema for `dns_ns_record_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'hits': Number of times the record has been used;



<a id="nestedblock--dns_ptr_record_list"></a>
### Nested Schema for `dns_ptr_record_list`

Required:

- `ptr_name` (String) Specify Domain Name

Optional:

- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--dns_ptr_record_list--sampling_enable))
- `ttl` (Number) Specify TTL
- `uuid` (String) uuid of the object

<a id="nestedblock--dns_ptr_record_list--sampling_enable"></a>
### Nested Schema for `dns_ptr_record_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'hits': Number of times the record has been used;



<a id="nestedblock--dns_record_list"></a>
### Nested Schema for `dns_record_list`

Required:

- `type` (Number) Specify DNS Type

Optional:

- `data` (String) Specify DNS Data
- `uuid` (String) uuid of the object


<a id="nestedblock--dns_srv_record_list"></a>
### Nested Schema for `dns_srv_record_list`

Required:

- `port` (Number) Specify Port (Port Number)
- `srv_name` (String) Specify Domain Name

Optional:

- `priority` (Number) Specify Priority
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--dns_srv_record_list--sampling_enable))
- `ttl` (Number) Specify TTL
- `uuid` (String) uuid of the object
- `weight` (Number) Specify Weight, default is 10

<a id="nestedblock--dns_srv_record_list--sampling_enable"></a>
### Nested Schema for `dns_srv_record_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'hits': Number of times the record has been used;



<a id="nestedblock--dns_txt_record_list"></a>
### Nested Schema for `dns_txt_record_list`

Required:

- `record_name` (String) Specify the Object Name for TXT Data

Optional:

- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--dns_txt_record_list--sampling_enable))
- `ttl` (Number) Specify TTL
- `txt_data` (String) Specify TXT Data
- `uuid` (String) uuid of the object

<a id="nestedblock--dns_txt_record_list--sampling_enable"></a>
### Nested Schema for `dns_txt_record_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'hits': Number of times the record has been used;



<a id="nestedblock--geo_location_list"></a>
### Nested Schema for `geo_location_list`

Required:

- `geo_name` (String) Specify the geo-location

Optional:

- `action` (Number) Action for this geo-location
- `action_type` (String) 'allow': Allow query from this geo-location; 'drop': Drop query from this geo-location; 'forward': Forward packet for this geo-location; 'ignore': Send empty response to this geo-location; 'reject': Send refuse response to this geo-location;
- `alias` (Block List) (see [below for nested schema](#nestedblock--geo_location_list--alias))
- `forward_type` (String) 'both': Forward both query and response; 'query': Forward query from this geo-location; 'response': Forward response to this geo-location;
- `policy` (String) Policy for this geo-location (Specify the policy name)
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--geo_location_list--alias"></a>
### Nested Schema for `geo_location_list.alias`

Optional:

- `alias` (String) Send CNAME response for this geo-location (Specify a CNAME record)



<a id="nestedblock--health_check_port"></a>
### Nested Schema for `health_check_port`

Optional:

- `health_check_port` (Number) Check Related Port Status (Port Number)


<a id="nestedblock--sampling_enable"></a>
### Nested Schema for `sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'received-query': Number of DNS queries received for the service; 'sent-response': Number of DNS replies sent to clients for the service; 'proxy-mode-response': Number of DNS replies sent to clients by the ACOS device as a DNS proxy for the service; 'cache-mode-response': Number of cached DNS replies sent to clients by the ACOS device for the service. (This statistic applies only if the DNS cache; 'server-mode-response': Number of DNS replies sent to clients by the ACOS device as a DNS server for the service. (This statistic applies only if the D; 'sticky-mode-response': Number of DNS replies sent to clients by the ACOS device to keep the clients on the same site. (This statistic applies only if; 'backup-mode-response': help Number of DNS replies sent to clients by the ACOS device in backup mode;

