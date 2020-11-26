---
layout: "thunder"
page_title: "thunder: thunder_ntp_server_hostname"
sidebar_current: "docs-thunder-resource-ntp-server-hostname"
description: |-
	Provides details about thunder ntp server hostname resource for A10
---

# thunder\_ntp\_server\_hostname

`thunder_ntp_server_hostname` Provides details about thunder ntp server hostname
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_ntp_server_hostname" "resourceNtpServerHostnameTest" {
	host_servername = "string"
key = 0
prefer = 0
action = "string"
uuid = "string"
 
}

```

## Argument Reference

* `action` - ‘enable’: Enable this NTP server; ‘disable’: Disable this NTP server;
* `host_servername` - IPV4 address, IPV6 address or host name of NTP server(string1~63)
* `key` - Use trusted key to authenticate this NTP server (The trusted key number to use)
* `prefer` - Set this NTP server as preferred
* `uuid` - uuid of the object

