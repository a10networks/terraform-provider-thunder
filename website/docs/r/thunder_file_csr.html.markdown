---
layout: "thunder"
page_title: "thunder: thunder_file_csr"
sidebar_current: "docs-thunder-resource-file-csr"
description: |-
    Provides details about thunder file csr resource for A10
---

# thunder\_file\_csr

`thunder_file_csr` Provides details about thunder file csr
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_file_csr" "resourceFileCsrTest" {
	file = "string"
size = 0
file_handle = "string"
action = "string"
dst_file = "string"
uuid = "string"
 
}

```

## Argument Reference

* `csr` - ssl certificate signing request file
* `file` - Specify the File Name
* `size` - CSR file size in byte
* `file-handle` - full path of the uploaded file
* `action` - 'export': export;
* `dst-file` - destination file name for copy and rename action
* `uuid` - uuid of the object

