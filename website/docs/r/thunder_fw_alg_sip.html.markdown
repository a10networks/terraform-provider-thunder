---
layout: "thunder"
page_title: "thunder: thunder_fw_alg_sip"
sidebar_current: "docs-thunder-resource-fw-alg-sip"
description: |-
	Provides details about thunder fw alg sip resource for A10
---

# thunder\_fw\_alg\_sip

`thunder_fw_alg_sip` Provides details about thunder fw alg sip
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_alg_sip" "Fw_Alg_Sip_Test" {

default_port_disable = "string"
sampling-enable {   
        counters1 =  "string" 
        }
uuid = "string"

}

```

## Argument Reference

* `default_port_disable` - ‘default-port-disable’: Disable SIP ALG default port 5060;
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘stat-request’: Request Received; ‘stat-response’: Response Received; ‘method-register’: Method REGISTER; ‘method-invite’: Method INVITE; ‘method-ack’: Method ACK; ‘method-cancel’: Method CANCEL; ‘method-bye’: Method BYE; ‘method-port-config’: Method OPTIONS; ‘method-prack’: Method PRACK; ‘method-subscribe’: Method SUBSCRIBE; ‘method-notify’: Method NOTIFY; ‘method-publish’: Method PUBLISH; ‘method-info’: Method INFO; ‘method-refer’: Method REFER; ‘method-message’: Method MESSAGE; ‘method-update’: Method UPDATE; ‘method-unknown’: Method Unknown; ‘parse-error’: Message Parse Error; ‘keep-alive’: Keep Alive; ‘contact-error’: Contact Process Error; ‘sdp-error’: SDP Process Error; ‘rtp-port-no-op’: RTP Port No Op; ‘rtp-rtcp-port-success’: RTP RTCP Port Success; ‘rtp-port-failure’: RTP Port Failure; ‘rtcp-port-failure’: RTCP Port Failure; ‘contact-port-no-op’: Contact Port No Op; ‘contact-port-success’: Contact Port Success; ‘contact-port-failure’: Contact Port Failure; ‘contact-new’: Contact Alloc; ‘contact-alloc-failure’: Contact Alloc Failure; ‘contact-eim’: Contact EIM; ‘contact-eim-set’: Contact EIM Set; ‘rtp-new’: RTP Alloc; ‘rtp-alloc-failure’: RTP Alloc Failure; ‘rtp-eim’: RTP EIM; ‘helper-found’: SMP Helper Conn Found; ‘helper-created’: SMP Helper Conn Created; ‘helper-deleted’: SMP Helper Conn Already Deleted; ‘helper-freed’: SMP Helper Conn Freed; ‘helper-failure’: SMP Helper Failure;

