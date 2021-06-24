---
layout: "thunder"
page_title: "thunder: thunder_glm_send"
sidebar_current: "docs-thunder-resource-glm-send"
description: |-
	Provides details about thunder glm send resource for A10
---

# thunder\_glm\_send

`thunder_glm_send` Provides details about thunder glm send
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_glm_send" "glm_send1" {

  depends_on = [thunder_ip_dns_primary.dns_primary, thunder_glm.glm1]
  license_request = 1

}
```

## Argument Reference

* `license_request` - Do a single GLM license retrieval

