---
layout: "thunder"
page_title: "thunder: thunder_fw_alg_rtsp"
sidebar_current: "docs-thunder-resource-fw-alg-rtsp"
description: |-
	Provides details about thunder fw alg rtsp resource for A10
---

# thunder\_fw\_alg\_rtsp

`thunder_fw_alg_rtsp` Provides details about thunder fw alg rtsp
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_alg_rtsp" "Fw_Alg_Rtsp_Test" {
default_port_disable = "string"
sampling-enable {   
        counters1 =  "string" 
        }
uuid = "string"
 
}

```

## Argument Reference

* `default_port_disable` - ‘default-port-disable’: Disable RTSP ALG default port 554;
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘transport-inserted’: Transport Created; ‘transport-freed’: Transport Freed; ‘transport-alloc-failure’: Transport Alloc Failure; ‘data-session-created’: Data Session Created; ‘data-session-freed’: Data Session Freed; ‘ext-creation-failure’: Extension Creation Failure; ‘transport-add-to-ext’: Transport Added to Extension; ‘transport-removed-from-ext’: Transport Removed from Extension; ‘transport-too-many’: Too Many Transports for Control Conn; ‘transport-already-in-ext’: Transport Already in Extension; ‘transport-exists’: Transport Already Exists; ‘transport-link-ext-failure-control’: Transport Link to Extension Failure Control; ‘transport-link-ext-data’: Transport Link to Extension Data; ‘transport-link-ext-failure-data’: Transport Link to Extension Failure Data; ‘transport-inserted-shadow’: Transport Inserted Shadow; ‘transport-creation-race’: Transport Create Race; ‘transport-alloc-failure-shadow’: Transport Alloc Failure Shadow; ‘transport-put-in-del-q’: Transport Put in Delete Queue; ‘transport-freed-shadow’: Transport Freed Shadow; ‘transport-acquired-from-control’: Transport Acquired Control; ‘transport-found-from-prev-control’: Transport Found From Prev Control; ‘transport-acquire-failure-from-control’: Transport Acquire Failure Control; ‘transport-released-from-control’: Transport Released Control; ‘transport-double-release-from-control’: Transport Double Release Control; ‘transport-acquired-from-data’: Transport Acquired Data; ‘transport-acquire-failure-from-data’: Transport Acquire Failure Data; ‘transport-released-from-data’: Transport Released Data; ‘transport-double-release-from-data’: Transport Double Release Data; ‘transport-retry-lookup-on-data-free’: Transport Retry Lookup Data; ‘transport-not-found-on-data-free’: Transport Not Found Data; ‘data-session-created-shadow’: Data Session Created Shadow; ‘data-session-freed-shadow’: Data Session Freed Shadow; ‘ha-control-ext-creation-failure’: HA Control Extension Creation Failure; ‘ha-control-session-created’: HA Control Session Created; ‘ha-data-session-created’: HA Data Session Created;

