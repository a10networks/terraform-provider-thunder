---
layout: "thunder"
page_title: "thunder: thunder_class_list"
sidebar_current: "docs-thunder-resource-class-list"
description: |-
    Provides details about thunder class list resource for A10
---

# thunder\_class\_list

`thunder_class_list` Provides details about thunder class list
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_class_list" "resourceClassListTest" {
	name = "string"
type = "string"
file = 0
ipv4-list {   
	ipv4addr =  "string" 
	lid =  0 
	glid =  0 
	shared_partition_glid =  0 
	glid_shared =  0 
	lsn_lid =  0 
	lsn_radius_profile =  0 
	gtp_rate_limit_policy_v4 =  "string" 
	age =  0 
	}
ipv6-list {   
	ipv6_addr =  "string" 
	v6_lid =  0 
	v6_glid =  0 
	shared_partition_v6_glid =  0 
	v6_glid_shared =  0 
	v6_lsn_lid =  0 
	v6_lsn_radius_profile =  0 
	gtp_rate_limit_policy_v6 =  "string" 
	v6_age =  0 
	}
dns {   
	dns_match_type =  "string" 
	dns_match_string =  "string" 
	dns_lid =  0 
	dns_glid =  0 
	shared_partition_dns_glid =  0 
	dns_glid_shared =  0 
	}
str-list {   
	str =  "string" 
	str_lid_dummy =  0 
	str_lid =  0 
	str_glid_dummy =  0 
	str_glid =  0 
	shared_partition_str_glid =  0 
	str_glid_shared =  0 
	value_str =  "string" 
	}
ac-list {   
	ac_match_type =  "string" 
	ac_key_string =  "string" 
	ac_value =  "string" 
	gtp_rate_limit_policy_str =  "string" 
	}
uuid = "string"
user_tag = "string"
 
}

```

## Argument Reference

* `class-list` - Configure classification list
* `name` - Specify name of the class list
* `type` - 'ac': Make class-list type Aho-Corasick; 'dns': Make class-list type DNS; 'ipv4': Make class-list type IPv4; 'ipv6': Make class-list type IPv6; 'string': Make class-list type String; 'string-case-insensitive': Make class-list type String-case-insensitive. Case insensitive is applied to key string;
* `file` - Create/Edit a class-list stored as a file
* `ipv4addr` - Specify IP address
* `lid` - Use Limit ID defined in template (Specify LID index)
* `glid` - Use global Limit ID (Specify global LID index)
* `shared-partition-glid` - Reference a glid from shared partition
* `glid-shared` - Use global Limit ID
* `lsn-lid` - LSN Limit ID (LID index)
* `lsn-radius-profile` - LSN RADIUS Profile Index
* `gtp-rate-limit-policy-v4` - GTP Rate Limit Template Name
* `age` - Specify age in minutes
* `ipv6-addr` - Specify IPv6 host or subnet
* `v6-lid` - Use Limit ID defined in template (Specify LID index)
* `v6-glid` - Use global Limit ID (Specify global LID index)
* `shared-partition-v6-glid` - Reference a glid from shared partition
* `v6-glid-shared` - Use global Limit ID
* `v6-lsn-lid` - LSN Limit ID (LID index)
* `v6-lsn-radius-profile` - LSN RADIUS Profile Index
* `gtp-rate-limit-policy-v6` - GTP Rate Limit Template Name
* `v6-age` - Specify age in minutes
* `dns-match-type` - 'contains': Domain contains another string; 'ends-with': Domain ends with another string; 'starts-with': Domain starts-with another string;
* `dns-match-string` - Domain name
* `dns-lid` - Use Limit ID defined in template (Specify LID index)
* `dns-glid` - Use global Limit ID (Specify global LID index)
* `shared-partition-dns-glid` - Reference a glid from shared partition
* `dns-glid-shared` - Use global Limit ID
* `str` - Specify key string
* `str-lid-dummy` - Use Limit ID defined in template
* `str-lid` - LID index
* `str-glid-dummy` - Use global Limit ID
* `str-glid` - Global LID index
* `shared-partition-str-glid` - Reference a glid from shared partition
* `str-glid-shared` - Use global Limit ID
* `value-str` - Specify value string
* `ac-match-type` - 'contains': String contains another string; 'ends-with': String ends with another string; 'equals': String equals another string; 'starts-with': String starts with another string;
* `ac-key-string` - Specify key string
* `ac-value` - Specify value string
* `gtp-rate-limit-policy-str` - GTP Rate Limit Template Name
* `uuid` - uuid of the object
* `user-tag` - Customized tag

