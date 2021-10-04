---
layout: "thunder"
page_title: "thunder: thunder_file_ssl_cert_key"
sidebar_current: "docs-thunder-resource-file-ssl-cert-key"
description: |-
    Provides details about thunder file ssl cert key resource for A10
---

# thunder\_file\_ssl\_cert\_key

`thunder_file_ssl_cert_key` Provides details about thunder file ssl cert key
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_file_ssl_cert_key" "resourceFileSslCertKeyTest" {
	file = "string"
size = 0
file_handle = "string"
secured = 0
file_local_path = "/path/to/pki-ssl-cert"
action = "string"
dst_file = "string"
 
}

```

## Argument Reference

* `ssl-cert-key` - ssl certificate and key file information and management commands
* `file` - ssl certificate local file name
* `size` - ssl certificate file size in byte
* `file-handle` - full path of the uploaded file
* `secured` - Mark keys as non-exportable
* `action` - 'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;
* `dst-file` - destination file name for copy and rename action
* `file_local_path` - local directory path for pki ssl-cert to upload

