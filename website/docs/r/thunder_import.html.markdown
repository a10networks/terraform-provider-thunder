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
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_import" "Import_Test" {

geo_location = "string"
ssl_cert_key = "string"
class_list_convert = "string"
bw_list = "string"
usb_license = "string"
ip_map_list = "string"
health_external {  
        description =  "string" 
        remote_file =  "string" 
        externalfilename =  "string" 
        password =  "string" 
        use_mgmt_port =  0 
        overwrite =  0 
        }
auth_portal = "string"
local_uri_file = "string"
aflex = "string"
overwrite = 0
class_list_type = "string"
pfx_password = "string"
web_category_license = "string"
thales_kmdata = "string"
secured = 0
ssl_crl = "string"
terminal = 0
policy = "string"
file_inspection_bw_list = "string"
thales_secworld = "string"
lw_4o6 = "string"
auth_portal_image = "string"
health_postfile {  
        postfilename =  "string" 
        password =  "string" 
        use_mgmt_port =  0 
        remote_file =  "string" 
        overwrite =  0 
        }
class_list = "string"
glm_license = "string"
dnssec_ds = "string"
cloud_creds = "string"
auth_jwks = "string"
wsdl = "string"
password = "string"
ssl_key = "string"
use_mgmt_port = 0
remote_file = "string"
cloud_config = "string"
to_device {  
        web_category_license =  "string" 
        remote_file =  "string" 
        glm_license =  "string" 
        glm_cert =  "string" 
        device =  0 
        use_mgmt_port =  0 
        overwrite =  0 
        }
user_tag = "string"
store_name = "string"
ca_cert = "string"
glm_cert = "string"
store {  
        create =  0 
        name =  "string" 
        remote_file =  "string" 
        delete =  0 
        }
xml_schema = "string"
certificate_type = "string"
auth_saml_idp {  
        remote_file =  "string" 
        saml_idp_name =  "string" 
        verify_xml_signature =  0 
        password =  "string" 
        use_mgmt_port =  0 
        overwrite =  0 
        }
ssl_cert = "string"
dnssec_dnskey = "string"
 
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

