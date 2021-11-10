---
layout: "thunder"
page_title: "thunder: thunder_slb_template_dns_class_list"
sidebar_current: "docs-thunder-resource-slb-template-dns-class-list"
description: |-
    Provides details about thunder slb template dns class list resource for A10
---

# thunder\_slb\_template\_dns\_class\_list

`thunder_slb_template_dns_class_list` Provides details about thunder slb template dns class list
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_dns_class_list" "resourceSlbTemplateDnsClassListTest" {
	name = "string"
uuid = "string"
lid-list {   
	lidnum =  0 
	conn_rate_limit =  0 
	per =  0 
	over_limit_action =  0 
	action_value =  "string" 
	lockout =  0 
	log =  0 
	log_interval =  0 
dns {  
 	cache_action =  "string" 
	ttl =  0 
	weight =  0 
	honor_server_response_ttl =  0 
	}
	uuid =  "string" 
	user_tag =  "string" 
	}
 
}

```

## Argument Reference

* `class-list` - Classification list
* `name` - Specify a class list name
* `uuid` - uuid of the object
* `lidnum` - Specify a limit ID
* `conn-rate-limit` - Connection rate limit
* `per` - Per (Number of 100ms)
* `over-limit-action` - Action when exceeds limit
* `action-value` - 'dns-cache-disable': Disable DNS cache when it exceeds limit; 'dns-cache-enable': Enable DNS cache when it exceeds limit; 'forward': Forward the traffic even it exceeds limit;
* `lockout` - Don't accept any new connection for certain time (Lockout duration in minutes)
* `log` - Log a message
* `log-interval` - Log interval (minute, by default system will log every over limit instance)
* `cache-action` - 'cache-disable': Disable dns cache; 'cache-enable': Enable dns cache;
* `ttl` - TTL for cache entry (TTL in seconds)
* `weight` - Weight for cache entry
* `honor-server-response-ttl` - Honor the server reponse TTL
* `user-tag` - Customized tag

