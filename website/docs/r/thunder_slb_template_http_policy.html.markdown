---
layout: "thunder"
page_title: "thunder: thunder_slb_template_http_policy"
sidebar_current: "docs-thunder-resource-slb-template-http-policy"
description: |-
	Provides details about thunder slb template http policy resource for A10
---

# thunder\_slb\_template\_http\_policy

`thunder_slb_template_http_policy` Provides details about thunder slb template http policy
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_http_policy" "Slb_Template_Http_Policy_Test" {

cookie_name = "string"
name = "string"
http-policy-match {   
        type =  "string" 
        match_string =  "string" 
        template_name =  "string" 
        service_group =  "string" 
        url_under_host =  "string" 
        other_match_string =  "string" 
        template =  "string" 
        other_match_type =  "string" 
        match_type =  "string" 
        }
uuid = "string"
user_tag = "string"
geo-location-match {   
        geo_location =  "string" 
        geo_location_service_group =  "string" 
        geo_location_template =  "string" 
        geo_location_template_name =  "string" 
        }
 
}

```

## Argument Reference

* `cookie_name` - name of cookie to match (Cookie Name)
* `name` - http-policy template name
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `match_string` - URL String
* `match_type` - ‘contains’: Select service group if URL string contains another string; ‘ends-with’: Select service group if URL string ends with another string; ‘equals’: Select service group if URL string equals another string; ‘starts-with’: Select service group if URL string starts with another string;
* `service_group` - Service Group to be used (Service Group Name)
* `template` - ‘waf’: waf;  (WAF template to be used)
* `template_name` - WAF template to be used (Template Name)
* `type` - ‘cookie’: cookie value match; ‘host’: hostname match; ‘url’: URL match;
* `geo_location` - Geolocation name
* `geo_location_service_group` - Service Group to be used (Service Group Name)
* `geo_location_template` - ‘waf’: waf;  (WAF template to be used)
* `geo_location_template_name` - WAF template to be used (Template Name)
