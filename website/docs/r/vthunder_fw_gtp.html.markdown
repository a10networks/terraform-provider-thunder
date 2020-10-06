---
layout: "vthunder"
page_title: "vthunder: vthunder_fw_gtp"
sidebar_current: "docs-vthunder-resource-fw-gtp"
description: |-
	Provides details about vthunder fw gtp resource for A10
---

# vthunder\_fw\_gtp

`vthunder_fw_gtp` Provides details about vthunder fw gtp
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_fw_gtp" "FwTest" {
	gtp_value = "enable" 
}
```

## Argument Reference

* `gtp_value` - ‘enable’: Enable GTP Inspection;
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘create-session-request’: Create Session Request; ‘create-session-response’: Create Session Response; ‘path-management-message’: Path Management Message; ‘delete-session-request’: Delete Session Request; ‘delete-session-response’: Delete Session Response; ‘reserved-field-set-drop’: Reserved field set drop; ‘tunnel-id-flag-drop’: Tunnel ID Flag Incorrect; ‘message-filtering-drop’: Message Filtering Drop; ‘reserved-information-element-drop’: Resevered Information Element Field Drop; ‘mandatory-information-element-drop’: Mandatory Information Element Field Drop; ‘filter-list-drop’: APN IMSI Information Filtering Drop; ‘invalid-teid-drop’: Invalid TEID Drop; ‘out-of-state-drop’: Out Of State Drop; ‘message-length-drop’: Message Length Exceeded; ‘unsupported-message-type-v2’: GTP v2 message type is not supported; ‘fast-conn-setup’: Fast Conn Setup Attempt; ‘out-of-session-memory’: Out of Session Memory; ‘no-fwd-route’: No Forward Route; ‘no-rev-route’: NO Reverse Route; ‘invalid-key’: Invalid TEID Field; ‘create-session-request-retransmit’: Retransmitted Create Session Request; ‘delete-session-request-retransmit’: Retransmitted Delete Session Request; ‘response-cause-not-accepted-drop’: Response Cause indicates Request not Accepted; ‘invalid-imsi-len-drop’: Invalid IMSI Length Drop; ‘invalid-apn-len-drop’: Invalid APN Length Drop; ‘create-pdp-context-request-v1’: GTP v1 Create PDP Context Request; ‘create-pdp-context-response-v1’: GTP v1 Create PDP Context Response; ‘path-management-message-v1’: GTP v1 Path Management Message; ‘reserved-field-set-drop-v1’: GTP v1 Reserved field set drop; ‘message-filtering-drop-v1’: GTP v1 Message Filtering Drop; ‘reserved-information-element-drop-v1’: GTP v1 Reserved Information Element Field Drop; ‘mandatory-information-element-drop-v1’: GTP v1 Mandatory Information Element Field Drop; ‘filter-list-drop-v1’: GTP v1 APN IMSI Information Filtering Drop; ‘invalid-teid-drop-v1’: GTP v1 Invalid TEID Drop; ‘message-length-drop-v1’: GTP v1 Message Length Exceeded; ‘version-not-supported’: GTP version is not supported; ‘unsupported-message-type-v1’: GTP v1 message type is not supported; ‘delete-pdp-context-request-v1’: GTP v1 Delete Context PDP Request; ‘delete-pdp-context-response-v1’: GTP v1 Delete Context PDP Response; ‘create-pdp-context-request-v0’: GTP v0 Create PDP Context Request; ‘create-pdp-context-response-v0’: GTP v0 Create PDP Context Response; ‘delete-pdp-context-request-v0’: GTP v0 Delete Context PDP Request; ‘delete-pdp-context-response-v0’: GTP v0 Delete Context PDP Response; ‘path-management-message-v0’: GTP v0 Path Management Message; ‘message-filtering-drop-v0’: GTP v0 Message Filtering Drop; ‘unsupported-message-type-v0’: GTP v0 message type is not supported; ‘invalid-flow-label-drop-v0’: GTP v0 Invalid flow label drop; ‘invalid-tid-drop-v0’: GTP v0 Invalid tid drop; ‘message-length-drop-v0’: GTP v0 Message Length Exceeded; ‘mandatory-information-element-drop-v0’: GTP v0 Mandatory Information Element Field Drop; ‘filter-list-drop-v0’: GTP v0 APN IMSI Information Filtering Drop; ‘gtp-in-gtp-drop’: GTP in GTP Filtering Drop;

