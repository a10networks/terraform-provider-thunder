---
layout: "thunder"
page_title: "thunder: thunder_slb_svm_source_nat"
sidebar_current: "docs-thunder-slb-svm-source-nat"
description: |-
    Provides details about thunder SLB svm source nat resource for A10
---

# thunder\_slb\_svm\_source\_nat

`thunder_slb_svm_source_nat` Provides details about thunder SLB svm source nat
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_svm_source_nat" "svm_source" {
	pool = "test" 
}
```

## Argument Reference

* `pool` - Specify NAT pool or pool group



