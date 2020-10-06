---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_cache"
sidebar_current: "docs-vthunder-resource-slb-template-cache"
description: |-
    Provides details about vthunder SLB template cache resource for A10
---

# vthunder\_slb\_template\_cache

`vthunder_slb_template_cache` Provides details about vthunder SLB template cache
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_cache" "cache" {
	name = "testcache"
	user_tag = "test_tag"
	accept_reload_req = 0
	default_policy_nocache = 0
	age = 4550
	disable_insert_via = 0
	replacement_policy = "LFU"
	disable_insert_age = 0
	max_content_size = 0
	max_cache_size = 700
	remove_cookies = 0
	verify_host = 0
	min_content_size = 0
	sampling_enable {
        counters1 = "all"
      }
}
```

## Argument Reference

* `name` - Specify cache template name
* `user_tag` - Customized tag
* `accept_reload_req` - Accept reload requests via cache-control directives in HTTP headers
* `default_policy_nocache` - Specify default policy to be to not cache
* `age` - Specify duration in seconds cached content valid, default is 3600 seconds (seconds that the cached content is valid (default 3600 seconds))
* `disable_insert_via` - Disable insertion of via header in response served from RAM cache
* `replacement_policy` - 'LFU’: LFU;
* `disable_insert_age` - Disable insertion of age header in response served from RAM cache
* `max_content_size` - Maximum size (bytes) of response that can be cached - default 81920 (80KB)
* `max_cache_size` - Specify maximum cache size in megabytes, default is 80MB (RAM cache size in megabytes (default 80MB))
* `remove_cookies` - Remove cookies in response and cache
* `verify_host` - Verify request using host before sending response from RAM cache
* `min_content_size` - Minimum size (bytes) of response that can be cached - default 512
* `counters1` - 'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;


