---
layout: "thunder"
page_title: "thunder: thunder_web_category_reputation_scope"
sidebar_current: "docs-thunder-resource-web-category-reputation-scope"
description: |-
    Provides details about thunder web category reputation scope resource for A10
---

# thunder\_web\_category\_reputation\_scope

`thunder_web_category_reputation_scope` Provides details about thunder web category reputation scope
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_web_category_reputation_scope" "resourceWebCategoryReputationScopeTest" {
	name = "string"
greater_than {  
 	greater_trustworthy =  0 
	greater_low_risk =  0 
	greater_moderate_risk =  0 
	greater_suspicious =  0 
	greater_malicious =  0 
	greater_threshold =  0 
	}
less_than {  
 	less_trustworthy =  0 
	less_low_risk =  0 
	less_moderate_risk =  0 
	less_suspicious =  0 
	less_malicious =  0 
	less_threshold =  0 
	}
uuid = "string"
user_tag = "string"
sampling-enable {   
	counters1 =  "string" 
	}
 
}

```

## Argument Reference

* `reputation-scope` - Configure the scope of reputation score
* `name` - Reputation Scope name
* `greater-trustworthy` - Reputation score is greater than or equal to 81
* `greater-low-risk` - Reputation score is greater than or equal to 61
* `greater-moderate-risk` - Reputation score is greater than or equal to 41
* `greater-suspicious` - Reputation score is greater than or equal to 21
* `greater-malicious` - Reputation score is greater than or equal to 1
* `greater-threshold` - Reputation score is greater than or equal to the customized score (1-100)
* `less-trustworthy` - Reputation score is less than or equal to 100
* `less-low-risk` - Reputation score is less than or equal to 80
* `less-moderate-risk` - Reputation score is less than or equal to 60
* `less-suspicious` - Reputation score is less than or equal to 40
* `less-malicious` - Reputation score is less than or equal to 20
* `less-threshold` - Reputation score is less than or equal to a customized value (1-100)
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `counters1` - 'all': all; 'trustworthy': Trustworthy level(81-100); 'low-risk': Low-risk level(61-80); 'moderate-risk': Moderate-risk level(41-60); 'suspicious': Suspicious level(21-40); 'malicious': Malicious level(1-20);

