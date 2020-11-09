---
layout: "thunder"
page_title: "thunder: thunder_fw_apply_changes"
sidebar_current: "docs-thunder-resource-fw-apply-changes"
description: |-
	Provides details about thunder fw apply changes resource for A10
---

# thunder\_fw\_apply\_changes

`thunder_fw_apply_changes` Provides details about thunder fw apply changes
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_apply_changes" "Fw_Apply_Changes_Test" {
        forced = 0
 
}
```

## Argument Reference

* `forced` - Force recompile rule-set
