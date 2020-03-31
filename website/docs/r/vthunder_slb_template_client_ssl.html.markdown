---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_template_client_ssl"
sidebar_current: "docs-vthunder-resource-slb-template-client-ssl"
description: |-
    Provides details about vthunder slb template client-ssl resource for A10
---

# vthunder\_slb\_template\_client\_ssl

`vthunder_slb_template_client_ssl` provides details about slb template client_ssl
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_template_client_ssl" "testname" {
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

