---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_dns_cache"
sidebar_current: "docs-vthunder-resource-slb-dns-cache"
description: |-
    Provides details about vthunder SLB dns-cache resource for A10
---

# vthunder\_slb\_dns\_cache

`vthunder_slb_dns_cache` Provides details about vthunder SLB dns-cache
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_dns_cache" "dns_cache" {
	sampling_enable {
		counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable` - 
	* `counters1` - 'all’: all; 'hits’: Cache hits; 'miss’: Cache misses; 'bytes_served’: Bytes served from cache; 'total_req’: Total requests received; 'caching_req’: Total requests to cache; 'nc_req_header’: nc_req_header; 'nc_res_header’: nc_res_header; 'rv_success’: rv_success; 'rv_failure’: rv_failure; 'ims_request’: ims_request; 'nm_response’: nm_response; 'rsp_type_CL’: rsp_type_CL; 'rsp_type_CE’: rsp_type_CE; 'rsp_type_304’: rsp_type_304; 'rsp_type_other’: rsp_type_other; 'rsp_no_compress’: rsp_no_compress; 'rsp_gzip’: rsp_gzip; 'rsp_deflate’: rsp_deflate; 'rsp_other’: rsp_other; 'nocache_match’: nocache_match; 'match’: match; 'invalidate_match’: invalidate_match; 'content_toobig’: content_toobig; 'content_toosmall’: content_toosmall; 'entry_create_failures’: entry_create_failures; 'mem_size’: mem_size; 'entry_num’: entry_num; 'replaced_entry’: replaced_entry; 'aging_entry’: aging_entry; 'cleaned_entry’: cleaned_entry;


