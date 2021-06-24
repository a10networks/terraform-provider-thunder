---
layout: "thunder"
page_title: "thunder: thunder_web_category_statistics"
sidebar_current: "docs-thunder-resource-web-category-statistics"
description: |-
    Provides details about thunder web category statistics resource for A10
---

# thunder\_web\_category\_statistics

`thunder_web_category_statistics` Provides details about thunder web category statistics
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_web_category_statistics" "resourceWebCategoryStatisticsTest" {
	uuid = "string"
sampling-enable {   
	counters1 =  "string" 
	}
 
}

```

## Argument Reference

* `statistics` - Statistics
* `uuid` - uuid of the object
* `counters1` - 'all': all; 'db-lookup': db-lookup; 'cloud-cache-lookup': cloud-cache-lookup; 'cloud-lookup': cloud-lookup; 'rtu-lookup': rtu-lookup; 'lookup-latency': lookup-latency; 'db-mem': db-mem; 'rtu-cache-mem': rtu-cache-mem; 'lookup-cache-mem': lookup-cache-mem;

