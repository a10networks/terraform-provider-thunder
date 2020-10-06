---
layout: "vthunder"
page_title: "vthunder: vthunder_import"
sidebar_current: "docs-vthunder-resource-import"
description: |-
    Provides details about vthunder import resource for A10
---

# vthunder\_import

`vthunder_import` provides details about Import
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_import" "import1" {
		password = ""
	    cloud_creds = "oci_api_key.pem"
	    use_mgmt_port = 1
	    overwrite = 1
	    remote_file = ""  
}
```

## Argument Reference

* `password` - password for the remote site
* `cloud_creds` - Cloud Credentials File
* `use_mgmt_port` - Use management port as source port
* `overwrite` - Overwrite existing file
* `remote_file` - profile name for remote url
