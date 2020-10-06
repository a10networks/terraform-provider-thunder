---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_svm_source_nat"
sidebar_current: "docs-vthunder-slb-svm-source-nat"
description: |-
    Provides details about vthunder SLB svm source nat resource for A10
---

# vthunder\_slb\_svm\_source\_nat

`vthunder_slb_svm_source_nat` Provides details about vthunder SLB svm source nat
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_svm_source_nat" "svm_source" {
	pool = "test" 
}
```

## Argument Reference

* `pool` - Specify NAT pool or pool group



