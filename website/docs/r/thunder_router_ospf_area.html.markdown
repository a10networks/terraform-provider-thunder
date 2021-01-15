---
layout: "thunder"
page_title: "thunder: thunder_router_ospf_area"
sidebar_current: "docs-thunder-resource-router-ospf-area"
description: |-
    Provides details about thunder router ospf area resource for A10
---

# thunder\_router\_ospf\_area

`thunder_router_ospf_area` Provides details about thunder router ospf area
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_ospf_area" "resourceRouterOspfAreaTest" {
	area_ipv4 = "string"
area_num = 0
auth_cfg {  
 	authentication =  0 
	message_digest =  0 
	}
filter-lists {   
	filter_list =  0 
	acl_name =  "string" 
	acl_direction =  "string" 
	plist_name =  "string" 
	plist_direction =  "string" 
	}
nssa_cfg {  
 	nssa =  0 
	no_redistribution =  0 
	no_summary =  0 
	translator_role =  "string" 
	default_information_originate =  0 
	metric =  0 
	metric_type =  0 
	}
default_cost = 0
range-list {   
	area_range_prefix =  "string" 
	option =  "string" 
	}
shortcut = "string"
stub_cfg {  
 	stub =  0 
	no_summary =  0 
	}
virtual-link-list {   
	virtual_link_ip_addr =  "string" 
	bfd =  0 
	hello_interval =  0 
	dead_interval =  0 
	retransmit_interval =  0 
	transmit_delay =  0 
	virtual_link_authentication =  0 
	virtual_link_auth_type =  "string" 
	authentication_key =  "string" 
	message_digest_key =  0 
	md5 =  "string" 
	}
uuid = "string"
 
}

```

## Argument Reference

* `area` - OSPF area parameters
* `area-ipv4` - OSPF area ID in IP address format
* `area-num` - OSPF area ID as a decimal value
* `authentication` - Enable authentication
* `message-digest` - Use message-digest authentication
* `filter-list` - Filter networks between OSPF areas
* `acl-name` - Filter networks by access-list (Name of an access-list)
* `acl-direction` - 'in': Filter networks sent to this area; 'out': Filter networks sent from this area;
* `plist-name` - Filter networks by prefix-list (Name of an IP prefix-list)
* `plist-direction` - 'in': Filter networks sent to this area; 'out': Filter networks sent from this area;
* `nssa` - Specify a NSSA area
* `no-redistribution` - No redistribution into this NSSA area
* `no-summary` - Do not inject inter-area routes into area
* `translator-role` - 'always': Translate always; 'candidate': Candidate for translator (default); 'never': Do not translate;
* `default-information-originate` - Originate Type 7 default into NSSA area
* `metric` - OSPF default metric (OSPF metric)
* `metric-type` - OSPF metric type (OSPF metric type for default routes)
* `default-cost` - Set the summary-default cost of a NSSA or stub area (Stub's advertised default summary cost)
* `area-range-prefix` - Area range for IPv4 prefix
* `option` - 'advertise': Advertise this range (default); 'not-advertise': DoNotAdvertise this range;
* `shortcut` - 'default': Set default shortcutting behavior; 'disable': Disable shortcutting through the area; 'enable': Enable shortcutting through the area;
* `stub` - Configure OSPF area as stub
* `virtual-link-ip-addr` - ID (IP addr) associated with virtual link neighbor
* `bfd` - Bidirectional Forwarding Detection (BFD)
* `hello-interval` - Hello packet interval (Seconds)
* `dead-interval` - Dead router detection time (Seconds)
* `retransmit-interval` - LSA retransmit interval (Seconds)
* `transmit-delay` - LSA transmission delay (Seconds)
* `virtual-link-authentication` - Enable authentication
* `virtual-link-auth-type` - 'message-digest': Use message-digest authentication; 'null': Use null authentication;
* `authentication-key` - Set authentication key (Authentication key (8 chars))
* `message-digest-key` - Set message digest key (Key ID)
* `md5` - Use MD5 algorithm (Authentication key (16 chars))
* `uuid` - uuid of the object

