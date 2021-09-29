---
layout: "thunder"
page_title: "thunder: thunder_file_ssl_key"
sidebar_current: "docs-thunder-resource-file-ssl-key"
description: |-
    Provides details about thunder file ssl key resource for A10
---

# thunder\_file\_ssl\_key

`thunder_file_ssl_key` Provides details about thunder file ssl key
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_file_ssl_key" "resourceFileSslKeyTest" {
	file = "string"
size = 0
file_handle = "string"
secured = 0
action = "string"
dst_file = "string"
uuid = "string"
 
}

```

## Argument Reference

* `ssl-key` - ssl key file information and management commands
* `file` - ssl key local file name
* `size` - ssl key file size in byte
* `file-handle` - full path of the uploaded file
* `secured` - Mark as non-exportable
* `action` - 'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;
* `dst-file` - destination file name for copy and rename action
* `uuid` - uuid of the object

