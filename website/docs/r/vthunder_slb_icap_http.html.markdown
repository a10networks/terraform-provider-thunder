---
layout: "thunder"
page_title: "thunder: thunder_slb_icap_http"
sidebar_current: "docs-thunder-resource-slb-icap_http"
description: |-
    Provides details about thunder SLB icap http resource for A10
---

# thunder\_slb\_icap\_http

`thunder_slb_icap_http` Provides details about thunder SLB icap http
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_icap_http" "icap" {
	sampling_enable {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'status_200’: Status code 200; 'status_201’: Status code 201; 'status_202’: Status code 202; 'status_203’: Status code 203; 'status_204’: Status code 204; 'status_205’: Status code 205; 'status_206’: Status code 206; 'status_207’: Status code 207; 'status_100’: Status code 100; 'status_101’: Status code 101; 'status_102’: Status code 102; 'status_300’: Status code 300; 'status_301’: Status code 301; 'status_302’: Status code 302; 'status_303’: Status code 303; 'status_304’: Status code 304; 'status_305’: Status code 305; 'status_306’: Status code 306; 'status_307’: Status code 307; 'status_400’: Status code 400; 'status_401’: Status code 401; 'status_402’: Status code 402; 'status_403’: Status code 403; 'status_404’: Status code 404; 'status_405’: Status code 405; 'status_406’: Status code 406; 'status_407’: Status code 407; 'status_408’: Status code 408; 'status_409’: Status code 409; 'status_410’: Status code 410; 'status_411’: Status code 411; 'status_412’: Status code 412; 'status_413’: Status code 413; 'status_414’: Status code 414; 'status_415’: Status code 415; 'status_416’: Status code 416; 'status_417’: Status code 417; 'status_418’: Status code 418; 'status_422’: Status code 422; 'status_423’: Status code 423; 'status_424’: Status code 424; 'status_425’: Status code 425; 'status_426’: Status code 426; 'status_449’: Status code 449; 'status_450’: Status code 450; 'status_500’: Status code 500; 'status_501’: Status code 501; 'status_502’: Status code 502; 'status_503’: Status code 503; 'status_504’: Status code 504; 'status_505’: Status code 505; 'status_506’: Status code 506; 'status_507’: Status code 507; 'status_508’: Status code 508; 'status_509’: Status code 509; 'status_510’: Status code 510; 'status_1xx’: status code 1XX; 'status_2xx’: status code 2XX; 'status_3xx’: status code 3XX; 'status_4xx’: status code 4XX; 'status_5xx’: status code 5XX; 'status_6xx’: status code 6XX; 'status_unknown’: Status code unknown;