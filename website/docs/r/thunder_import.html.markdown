---
layout: "thunder"
page_title: "thunder: thunder_import"
sidebar_current: "docs-thunder-resource-import"
description: |-
	Provides details about thunder import resource for A10
---

# thunder\_import

`thunder_import` Provides details about thunder import
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_import" "test_import_aflex" {
      
    remote_file = "scp://root@10.1.1.1:/home/vip-switch-ar"
    use_mgmt_port =  1
    aflex =  "service-group-name"
    password = "password"
    overwrite =  1
}

```

## Argument Reference

* `aflex` - aFleX Script Source File
* `auth_jwks` - JSON web key
* `auth_portal` - Portal file for http authentication
* `auth_portal_image` - Image file for default portal
* `bw_list` - Black white List File
* `ca_cert` - CA Cert File(enter bulk when import an archive file)
* `certificate_type` - ‘pem’: pem; ‘der’: der; ‘pfx’: pfx; ‘p7b’: p7b;
* `class_list` - Class List File
* `class_list_convert` - Convert Class List File to A10 format
* `class_list_type` - ‘ac’: ac; ‘ipv4’: ipv4; ‘ipv6’: ipv6; ‘string’: string; ‘string-case-insensitive’: string-case-insensitive;
* `dnssec_dnskey` - DNSSEC DNSKEY(KSK) file for child zone
* `dnssec_ds` - DNSSEC DS file for child zone
* `file_inspection_bw_list` - Black white List File
* `geo_location` - Geo-location CSV File
* `glm_cert` - GLM certificate
* `glm_license` - License File
* `ip_map_list` - IP Map List File
* `local_uri_file` - Local URI files for http response
* `lw_4o6` - LW-4over6 Binding Table File
* `overwrite` - Overwrite existing file
* `password` - password for the remote site
* `pfx_password` - The password for certificate file (pfx type only)
* `policy` - WAF policy File
* `remote_file` - Profile name for remote url
* `ssl_cert` - SSL Cert File(enter bulk when import an archive file)
* `ssl_cert_key` - ‘bulk’: import an archive file;
* `ssl_crl` - SSL Crl File
* `ssl_key` - SSL Key File(enter bulk when import an archive file)
* `store_name` - Import store name
* `terminal` - terminal vi
* `thales_kmdata` - Thales Kmdata files
* `thales_secworld` - Thales security world files
* `usb_license` - USB License File
* `use_mgmt_port` - Use management port as source port
* `user_tag` - Customized tag
* `web_category_license` - License file to enable web-category feature
* `wsdl` - Web Service Definition Language File
* `xml_schema` - XML-Schema File
* `description` - Describe the Program Function briefly
* `externalfilename` - Specify the Program Name
* `postfilename` - Specify the File Name
* `create` - Create an import store profile
* `delete` - Delete an import store profile
* `name` - profile name to store remote url
* `device` - Device (Device ID)
* `saml_idp_name` - Metadata name
* `verify_xml_signature` - Verify metadata’s XML signature

