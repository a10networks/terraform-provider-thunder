---
layout: "thunder"
page_title: "thunder: thunder_slb_template_cache"
sidebar_current: "docs-thunder-resource-slb-template-cache"
description: |-
	Provides details about thunder slb template cache resource for A10
---

# thunder\_slb\_template\_cache

`thunder_slb_template_cache` Provides details about thunder slb template cache
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_slb_template_cache" "Slb_Template_Cache_Test" {

accept_reload_req = 0
name = "string"
default_policy_nocache = 0
age = 0
disable_insert_via = 0
user_tag = "string"
local-uri-policy {   
        local_uri =  "string" 
        }
sampling-enable {   
        counters1 =  "string" 
        }
replacement_policy = "string"
disable_insert_age = 0
max_content_size = 0
max_cache_size = 0
logging = "string"
uri-policy {   
        cache_action =  "string" 
        cache_value =  0 
        uri =  "string" 
        invalidate =  "string" 
        }
remove_cookies = 0
verify_host = 0
min_content_size = 0
uuid = "string"
 
}

```

## Argument Reference

* `accept_reload_req` - Accept reload requests via cache-control directives in HTTP headers
* `age` - Specify duration in seconds cached content valid, default is 3600 seconds (seconds that the cached content is valid (default 3600 seconds))
* `default_policy_nocache` - Specify default policy to be to not cache
* `disable_insert_age` - Disable insertion of age header in response served from RAM cache
* `disable_insert_via` - Disable insertion of via header in response served from RAM cache
* `logging` - Specify logging template (Logging Config name)
* `max_cache_size` - Specify maximum cache size in megabytes, default is 80MB (RAM cache size in megabytes (default 80MB))
* `max_content_size` - Maximum size (bytes) of response that can be cached - default 81920 (80KB)
* `min_content_size` - Minimum size (bytes) of response that can be cached - default 512
* `name` - Specify cache template name
* `remove_cookies` - Remove cookies in response and cache
* `replacement_policy` - ‘LFU’: LFU;
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `verify_host` - Verify request using host before sending response from RAM cache
* `local_uri` - Specify Local URI for caching (Specify URI pattern that the policy should be applied to, maximum 63 charaters)
* `counters1` - ‘all’: all; ‘hits’: hits; ‘miss’: miss; ‘bytes_served’: bytes_served; ‘total_req’: total_req; ‘caching_req’: caching_req; ‘nc_req_header’: nc_req_header; ‘nc_res_header’: nc_res_header; ‘rv_success’: rv_success; ‘rv_failure’: rv_failure; ‘ims_request’: ims_request; ‘nm_response’: nm_response; ‘rsp_type_CL’: rsp_type_CL; ‘rsp_type_CE’: rsp_type_CE; ‘rsp_type_304’: rsp_type_304; ‘rsp_type_other’: rsp_type_other; ‘rsp_no_compress’: rsp_no_compress; ‘rsp_gzip’: rsp_gzip; ‘rsp_deflate’: rsp_deflate; ‘rsp_other’: rsp_other; ‘nocache_match’: nocache_match; ‘match’: match; ‘invalidate_match’: invalidate_match; ‘content_toobig’: content_toobig; ‘content_toosmall’: content_toosmall; ‘entry_create_failures’: entry_create_failures; ‘mem_size’: mem_size; ‘entry_num’: entry_num; ‘replaced_entry’: replaced_entry; ‘aging_entry’: aging_entry; ‘cleaned_entry’: cleaned_entry;
* `cache_action` - ‘cache’: Specify if certain URIs should be cached; ‘nocache’: Specify if certain URIs should not be cached;
* `cache_value` - Specify seconds that content should be cached, default is age specified in cache template
* `invalidate` - Specify if URI should invalidate cache entries matching pattern (pattern that would match entries to be invalidated (64 chars max))
* `uri` - Specify URI for cache policy (Specify URI pattern that the policy should be applied to, maximum 63 charaters)
