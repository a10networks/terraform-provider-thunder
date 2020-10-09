---
layout: "thunder"
page_title: "thunder: thunder_import"
sidebar_current: "docs-thunder-resource-import"
description: |-
    Provides details about thunder import resource for A10
---

# thunder\_import

`thunder_import` provides details about Import
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_import" "import1" {
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
