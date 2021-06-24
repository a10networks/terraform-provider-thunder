---
layout: "thunder"
page_title: "thunder: thunder_web_category_proxy_server"
sidebar_current: "docs-thunder-resource-web-category-proxy-server"
description: |-
    Provides details about thunder web category proxy server resource for A10
---

# thunder\_web\_category\_proxy\_server

`thunder_web_category_proxy_server` Provides details about thunder web category proxy server
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_web_category_proxy_server" "resourceWebCategoryProxyServerTest" {
	proxy_host = "string"
http_port = 0
https_port = 0
auth_type = "string"
domain = "string"
username = "string"
password = 0
secret_string = "string"
uuid = "string"
 
}

```

## Argument Reference

* `proxy-server` - Commands to connect web-category through proxy server
* `proxy-host` - Proxy server hostname or IP address
* `http-port` - Proxy server HTTP port
* `https-port` - Proxy server HTTPS port(HTTP port will be used if not configured)
* `auth-type` - 'ntlm': NTLM authentication(default); 'basic': Basic authentication;
* `domain` - Realm for NTLM authentication
* `username` - Username for proxy authentication
* `password` - Password for proxy authentication
* `secret-string` - password value
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)
* `uuid` - uuid of the object

