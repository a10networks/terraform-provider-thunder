---
layout: "thunder"
page_title: "thunder: thunder_fw_alg_ftp"
sidebar_current: "docs-thunder-resource-fw-alg-ftp"
description: |-
	Provides details about thunder fw alg ftp resource for A10
---

# thunder\_fw\_alg\_ftp

`thunder_fw_alg_ftp` Provides details about thunder fw alg ftp
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_alg_ftp" "Fw_Alg_Ftp_Test" {
default_port_disable = "string"
sampling-enable {   
counters1 =  "string" 
        }
uuid = "string"
 
}


```

## Argument Reference

* `default_port_disable` - ‘default-port-disable’: Disable FTP ALG default port 21;
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘client-port-request’: PORT Requests From Client; ‘client-eprt-request’: EPRT Requests From Client; ‘server-pasv-reply’: PASV Replies From Server; ‘server-epsv-reply’: EPSV Replies From Server; ‘port-retransmits’: PORT Retransmits; ‘pasv-retransmits’: PASV Retransmits; ‘smp-app-type-mismatch’: SMP App Type Mismatch; ‘retransmit-sanity-check-failure’: Retransmit Sanity Check Failure; ‘smp-conn-alloc-failure’: SMP Helper Conn Alloc Failure; ‘port-helper-created’: PORT Helper Created; ‘pasv-helper-created’: PASV Helper Created; ‘port-helper-acquire-in-del-q’: PORT Helper Acquire In Del Queue; ‘port-helper-acquire-already-used’: PORT Helper Acquire Already Used; ‘pasv-helper-acquire-in-del-q’: PASV Helper Acquire In Del Queue; ‘pasv-helper-acquire-already-used’: PASV Helper Acquire Already Used; ‘port-helper-freed-used’: PORT Helper Freed Used; ‘port-helper-freed-unused’: PORT Helper Freed Unused; ‘pasv-helper-freed-used’: PASV Helper Freed Used; ‘pasv-helper-freed-unused’: PASV Helper Freed Unused;

