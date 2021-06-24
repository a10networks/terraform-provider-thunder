---
layout: "thunder"
page_title: "thunder: thunder_fw_logging"
sidebar_current: "docs-thunder-resource-fw-logging"
description: |-
	Provides details about thunder fw logging resource for A10
---

# thunder\_fw\_logging

`thunder_fw_logging` Provides details about thunder fw logging
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_fw_logging" "Fw_Logging_Test" {
        gtp {  
 sampling-enable {   
        counters1 =  "string" 
        }
        uuid =  "string" 
        }
sampling-enable {   
        counters1 =  "string" 
        }
name = "string"
uuid = "string"
 
}

```

## Argument Reference

* `name` - Logging Template Name
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘log_message_sent’: Log Packet Sent; ‘log_type_reset’: Log Event Type Reset; ‘log_type_deny’: Log Event Type Deny; ‘log_type_session_closed’: Log Event Type Session Close; ‘log_type_session_opened’: Log Event Type Session Open; ‘rule_not_logged’: Firewall Rule Not Logged; ‘log-dropped’: Log Packets Dropped; ‘tcp-session-created’: TCP Session Created; ‘tcp-session-deleted’: TCP Session Deleted; ‘udp-session-created’: UDP Session Created; ‘udp-session-deleted’: UDP Session Deleted; ‘icmp-session-deleted’: ICMP Session Deleted; ‘icmp-session-created’: ICMP Session Created; ‘icmpv6-session-deleted’: ICMPV6 Session Deleted; ‘icmpv6-session-created’: ICMPV6 Session Created; ‘other-session-deleted’: Other Session Deleted; ‘other-session-created’: Other Session Created; ‘http-request-logged’: HTTP Request Logged; ‘http-logging-invalid-format’: HTTP Logging Invalid Format Error; ‘dcmsg_permit’: Dcmsg Permit; ‘alg_override_permit’: Alg Override Permit; ‘template_error’: Template Error; ‘ipv4-frag-applied’: IPv4 Fragmentation Applied; ‘ipv4-frag-failed’: IPv4 Fragmentation Failed; ‘ipv6-frag-applied’: IPv6 Fragmentation Applied; ‘ipv6-frag-failed’: IPv6 Fragmentation Failed; ‘out-of-buffers’: Out of Buffers; ‘add-msg-failed’: Add Message to Buffer Failed; ‘tcp-logging-conn-established’: TCP Logging Conn Established; ‘tcp-logging-conn-create-failed’: TCP Logging Conn Create Failed; ‘tcp-logging-conn-dropped’: TCP Logging Conn Dropped; ‘log-message-too-long’: Log message too long; ‘http-out-of-order-dropped’: HTTP out-of-order dropped; ‘http-alloc-failed’: HTTP Request Info Allocation Failed; ‘sctp-session-created’: SCTP Session Created; ‘sctp-session-deleted’: SCTP Session Deleted; ‘log_type_sctp_inner_proto_filter’: Log Event Type SCTP Inner Proto Filter; ‘log_type_gtp_message_filtering’: Log Event Type GTP Message Filtering; ‘log_type_gtp_apn_filtering’: Log Event Type GTP Apn Filtering; ‘tcp-logging-port-allocated’: TCP Logging Port Allocated; ‘tcp-logging-port-freed’: TCP Logging Port Freed; ‘tcp-logging-port-allocation-failed’: TCP Logging Port Allocation Failed; ‘log_type_gtp_invalid_teid’: Log Event Type GTP Invalid TEID; ‘log_gtp_type_reserved_ie_present’: Log Event Type GTP Reserved Information Element Present; ‘log_type_gtp_mandatory_ie_missing’: Log Event Type GTP Mandatory Information Element Missing;
