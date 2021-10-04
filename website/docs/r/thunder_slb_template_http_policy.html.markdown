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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_http_policy" "resourceSlbTemplateHttpPolicyTest" {
	name = "string"
cookie_name = "string"
http-policy-match {   
	type =  "string" 
	match_type =  "string" 
	match_string =  "string" 
	service_group =  "string" 
	template =  "string" 
	template_name =  "string" 
	}
geo-location-match {   
	geo_location =  "string" 
	geo_location_service_group =  "string" 
	geo_location_template =  "string" 
	geo_location_template_name =  "string" 
	}
uuid = "string"
user_tag = "string"
multi-match-rule-list {   
	multi_match =  "string" 
	seq_num =  0 
	host_equals_type =  "string" 
	host_equals_string =  "string" 
	host_contains_type =  "string" 
	host_contains_string =  "string" 
	host_starts_with_type =  "string" 
	host_starts_with_string =  "string" 
	host_ends_with_type =  "string" 
	host_ends_with_string =  "string" 
	cookie_name_equals_type =  "string" 
	cookie_name_equals_string =  "string" 
	cookie_name_contains_type =  "string" 
	cookie_name_contains_string =  "string" 
	cookie_name_starts_with_type =  "string" 
	cookie_name_starts_with_string =  "string" 
	cookie_name_ends_with_type =  "string" 
	cookie_name_ends_with_string =  "string" 
	cookie_value_equals_type =  "string" 
	cookie_value_equals_string =  "string" 
	cookie_value_contains_type =  "string" 
	cookie_value_contains_string =  "string" 
	cookie_value_starts_with_type =  "string" 
	cookie_value_starts_with_string =  "string" 
	cookie_value_ends_with_type =  "string" 
	cookie_value_ends_with_string =  "string" 
	url_equals_type =  "string" 
	url_equals_string =  "string" 
	url_contains_type =  "string" 
	url_contains_string =  "string" 
	url_starts_with_type =  "string" 
	url_starts_with_string =  "string" 
	url_ends_with_type =  "string" 
	url_ends_with_string =  "string" 
	header_name_equals_type =  "string" 
	header_name_equals_string =  "string" 
	header_name_contains_type =  "string" 
	header_name_contains_string =  "string" 
	header_name_starts_with_type =  "string" 
	header_name_starts_with_string =  "string" 
	header_name_ends_with_type =  "string" 
	header_name_ends_with_string =  "string" 
	header_value_equals_type =  "string" 
	header_value_equals_string =  "string" 
	header_value_contains_type =  "string" 
	header_value_contains_string =  "string" 
	header_value_starts_with_type =  "string" 
	header_value_starts_with_string =  "string" 
	header_value_ends_with_type =  "string" 
	header_value_ends_with_string =  "string" 
	query_param_name_equals_type =  "string" 
	query_param_name_equals_string =  "string" 
	query_param_name_contains_type =  "string" 
	query_param_name_contains_string =  "string" 
	query_param_name_starts_with_type =  "string" 
	query_param_name_starts_with_string =  "string" 
	query_param_name_ends_with_type =  "string" 
	query_param_name_ends_with_string =  "string" 
	query_param_value_equals_type =  "string" 
	query_param_value_equals_string =  "string" 
	query_param_value_contains_type =  "string" 
	query_param_value_contains_string =  "string" 
	query_param_value_starts_with_type =  "string" 
	query_param_value_starts_with_string =  "string" 
	query_param_value_ends_with_type =  "string" 
	query_param_value_ends_with_string =  "string" 
	service_group =  "string" 
	template_waf =  "string" 
	uuid =  "string" 
	user_tag =  "string" 
	}
 
}

```

## Argument Reference

* `http-policy` - http-policy template
* `name` - http-policy template name
* `cookie-name` - name of cookie to match (Cookie Name)
* `type` - 'cookie': cookie value match; 'host': hostname match; 'url': URL match; 'header-name': header name match; 'header-value': header value match; 'query-param-name': query parameter name; 'query-param-value': query parameter value;
* `match-type` - 'contains': Select service group if URL string contains another string; 'ends-with': Select service group if URL string ends with another string; 'equals': Select service group if URL string equals another string; 'starts-with': Select service group if URL string starts with another string;
* `match-string` - URL String, use "[no-name]" for empty query-param-name match, use "[no-value]" for empty query-param-value match
* `service-group` - Service Group to be used (Service Group Name)
* `template` - 'waf': waf;  (WAF template to be used)
* `template-name` - WAF template to be used (Template Name)
* `geo-location` - Geolocation name
* `geo-location-service-group` - Service Group to be used (Service Group Name)
* `geo-location-template` - 'waf': waf;  (WAF template to be used)
* `geo-location-template-name` - WAF template to be used (Template Name)
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `multi-match` - Specify a multi-match-rule name
* `seq-num` - Specify a sequence number
* `host-equals-type` - 'equals': Host equals to string;
* `host-equals-string` - Host string
* `host-contains-type` - 'contains': Host contains string;
* `host-contains-string` - Host string
* `host-starts-with-type` - 'starts-with': Host starts-with string;
* `host-starts-with-string` - Host string
* `host-ends-with-type` - 'ends-with': Host ends-with string;
* `host-ends-with-string` - Host string
* `cookie-name-equals-type` - 'equals': Cookie name equals to string;
* `cookie-name-equals-string` - Cookie name string
* `cookie-name-contains-type` - 'contains': Cookie name contains string;
* `cookie-name-contains-string` - Cookie value string
* `cookie-name-starts-with-type` - 'starts-with': Cookie name starts-with string;
* `cookie-name-starts-with-string` - Cookie name string
* `cookie-name-ends-with-type` - 'ends-with': Cookie name ends-with string;
* `cookie-name-ends-with-string` - Cookie name string
* `cookie-value-equals-type` - 'equals': Cookie value equals to string;
* `cookie-value-equals-string` - Cookie value string
* `cookie-value-contains-type` - 'contains': Cookie value contains string;
* `cookie-value-contains-string` - Cookie value string
* `cookie-value-starts-with-type` - 'starts-with': Cookie value starts-with string;
* `cookie-value-starts-with-string` - Cookie value string
* `cookie-value-ends-with-type` - 'ends-with': Cookie value ends-with string;
* `cookie-value-ends-with-string` - Cookie value string
* `url-equals-type` - 'equals': URL equals to string;
* `url-equals-string` - URL string
* `url-contains-type` - 'contains': URL contains string;
* `url-contains-string` - URL string
* `url-starts-with-type` - 'starts-with': URL starts-with string;
* `url-starts-with-string` - URL string
* `url-ends-with-type` - 'ends-with': URL ends-with string;
* `url-ends-with-string` - URL string
* `header-name-equals-type` - 'equals': Header name equals to string;
* `header-name-equals-string` - Header name string
* `header-name-contains-type` - 'contains': Header name contains string;
* `header-name-contains-string` - Header name string
* `header-name-starts-with-type` - 'starts-with': Header name starts-with string;
* `header-name-starts-with-string` - Header name string
* `header-name-ends-with-type` - 'ends-with': Header name ends-with string;
* `header-name-ends-with-string` - Header name string
* `header-value-equals-type` - 'equals': Header value equals to string;
* `header-value-equals-string` - Header value string
* `header-value-contains-type` - 'contains': Header value contains string;
* `header-value-contains-string` - Header value string
* `header-value-starts-with-type` - 'starts-with': Header value starts-with string;
* `header-value-starts-with-string` - Header value string
* `header-value-ends-with-type` - 'ends-with': Header value ends-with string;
* `header-value-ends-with-string` - Header value string
* `query-param-name-equals-type` - 'equals': query parameter name equals to string;
* `query-param-name-equals-string` - query parameter name string, use "[no-name]" for empty query-param-name match
* `query-param-name-contains-type` - 'contains': query parameter name contains string;
* `query-param-name-contains-string` - query parameter name string
* `query-param-name-starts-with-type` - 'starts-with': query parameter name starts-with string;
* `query-param-name-starts-with-string` - query parameter name string
* `query-param-name-ends-with-type` - 'ends-with': query parameter name ends-with string;
* `query-param-name-ends-with-string` - query parameter name string
* `query-param-value-equals-type` - 'equals': query parameter value equals to string;
* `query-param-value-equals-string` - query parameter value string, use "[no-value]" for empty query-param-value match
* `query-param-value-contains-type` - 'contains': query parameter value contains string;
* `query-param-value-contains-string` - query parameter value string
* `query-param-value-starts-with-type` - 'starts-with': query parameter value starts-with string;
* `query-param-value-starts-with-string` - query parameter value string
* `query-param-value-ends-with-type` - 'ends-with': query parameter value ends-with string;
* `query-param-value-ends-with-string` - query parameter value string
* `template-waf` - Waf Template to be used (Waf Template Name)

