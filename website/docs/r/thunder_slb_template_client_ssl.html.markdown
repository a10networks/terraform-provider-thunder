---
layout: "thunder"
page_title: "thunder: thunder_slb_template_client_ssl"
sidebar_current: "docs-thunder-resource-slb-template-client-ssl"
description: |-
    Provides details about thunder slb template client-ssl resource for A10
---

# thunder\_slb\_template\_client\_ssl

`thunder_slb_template_client_ssl` provides details about slb template client_ssl
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_client_ssl" "client_ssl" {
	name = "testclientssl"
	user_tag = "test_tag"
	forward_proxy_ssl_version = 33
	forward_proxy_crl_disable = 1
	fp_cert_ext_aia_ocsp = "ocp" 
}
```

## Argument Reference

* `name` - Client SSL Template Name
* `user_tag` - Customized tag
* `forward_proxy_ssl_version` - TLS/SSL version, default is TLS1.2 (TLS/SSL version: 31-TLSv1.0, 32-TLSv1.1 and 33-TLSv1.2)
* `forward_proxy_crl_disable` - Disable Certificate Revocation List checking for forward proxy
* `fp_cert_ext_aia_ocsp` - OCSP (Authority Information Access URI)
