---
layout: "thunder"
page_title: "thunder: thunder_file_ssl_cert"
sidebar_current: "docs-thunder-resource-file-ssl-cert"
description: |-
    Provides details about thunder file ssl cert resource for A10
---

# thunder\_file\_ssl\_cert

`thunder_file_ssl_cert` Provides details about thunder file ssl cert
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_file_ssl_cert" "resourceFileSslCertTest" {
	file = "string"
size = 0
file_handle = "string"
certificate_type = "string"
pfx_password = "string"
file_local_path = "/path/to/pki-ssl-cert"
pfx_password_export = "string"
action = "string"
dst_file = "string"
uuid = "string"
 
}

```

## Argument Reference

* `ssl-cert` - ssl certificate file information and management commands
* `file` - ssl certificate local file name
* `size` - ssl certificate file size in byte
* `file-handle` - full path of the uploaded file
* `certificate-type` - 'pem': pem; 'der': der; 'pfx': pfx; 'p7b': p7b;
* `pfx-password` - The password for certificate file (pfx type only)
* `pfx-password-export` - The password for exported certificate file (pfx type only)
* `action` - 'create': create; 'import': import; 'export': export; 'copy': copy; 'rename': rename; 'check': check; 'replace': replace; 'delete': delete;
* `dst-file` - destination file name for copy and rename action
* `uuid` - uuid of the object
* `file_local_path` - local directory path for pki ssl-cert to upload
