---
layout: "thunder"
page_title: "thunder: thunder_glm"
sidebar_current: "docs-thunder-resource-glm"
description: |-
	Provides details about thunder glm resource for A10
---

# thunder\_glm

`thunder_glm` Provides details about thunder glm
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_glm" "Glm_Test" {
uuid = "string"
use_mgmt_port = 0
burst = 0
interval = 0
send {  
        license_request =  0 
        }
token = "string"
enterprise = "string"
proxy_server {  
        username =  "string" 
        uuid =  "string" 
        encrypted =  "Unknown Type: encrypted" 
        host =  "string" 
        password =  0 
        port =  0 
        secret_string =  "string" 
        }
appliance_name = "string"
enable_requests = 0
allocate_bandwidth = 0
port = 0
 
}

```

## Argument Reference

* `allocate_bandwidth` - Enter the requested bandwidth in Mbps for Capacity Pool
* `appliance_name` - Helpful identifier for this appliance
* `enable_requests` - Enable connection to the GLM
* `enterprise` - Enter the ELM hostname or IP
* `interval` - Enter the desired ping interval in hours
* `port` - Proxy server port
* `token` - License entitlement token
* `use_mgmt_port` - Use management port to connect to GLM
* `uuid` - uuid of the object
* `license_request` - Do a single GLM license retrieval
* `encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED secret string)
* `host` - Proxy server hostname or IP address
* `password` - Password for proxy authentication
* `secret_string` - password value
* `username` - Username for proxy authentication
