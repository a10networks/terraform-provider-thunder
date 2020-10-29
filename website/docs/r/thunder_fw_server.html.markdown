---
layout: "thunder"
page_title: "thunder: thunder_fw_server"
sidebar_current: "docs-thunder-resource-fw-server"
description: |-
	Provides details about thunder fw server resource for A10
---

# thunder\_fw\_server

`thunder_fw_server` Provides details about thunder fw server
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_server" "Fw_Server_Test" {

health_check_disable = 0
port-list {   
        health_check_disable =  0 
        protocol =  "string" 
        uuid =  "string" 
        user_tag =  "string" 
sampling-enable {   
        counters1 =  "string" 
        }
        port_number =  0 
        action =  "string" 
        health_check =  "string" 
        }
uuid = "string"
fqdn_name = "string"
resolve_as = "string"
sampling-enable {   
        counters1 =  "string" 
        }
user_tag = "string"
host = "string"
action = "string"
server_ipv6_addr = "string"
health_check = "string"
name = "string"
 
}
```

## Argument Reference

* `action` - ‘enable’: enable; ‘disable’: disable;
* `fqdn_name` - Server hostname
* `health_check` - Health Check (Monitor Name)
* `health_check_disable` - Disable health check
* `host` - IP Address
* `name` - Server Name
* `resolve_as` - ‘resolve-to-ipv4’: Use A Query only to resolve FQDN; ‘resolve-to-ipv6’: Use AAAA Query only to resolve FQDN; ‘resolve-to-ipv4-and-ipv6’: Use A as well as AAAA Query to resolve FQDN;
* `server_ipv6_addr` - IPV6 address
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `port_number` - Port Number
* `protocol` - ‘tcp’: TCP Port; ‘udp’: UDP Port;
* `counters1` - ‘all’: all; ‘curr-conn’: Current connections; ‘total-conn’: Total connections; ‘fwd-pkt’: Forward packets; ‘rev-pkt’: Reverse Packets; ‘peak-conn’: Peak connections;

