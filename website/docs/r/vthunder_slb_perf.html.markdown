---
layout: "vthunder"
page_title: "vthunder: vthunder_slb_perf"
sidebar_current: "docs-vthunder-resource-slb-perf"
description: |-
    Provides details about vthunder SLB perf resource for A10
---

# vthunder\_slb\_perf

`vthunder_slb_perf` Provides details about vthunder SLB perf
## Example Usage


```hcl
provider "vthunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_slb_perf" "perf" {
	sampling_enable  {
	    counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable`
    * `counters1` - 'all’: all; 'total-throughput-bits-per-sec’: Total Throughput in bits/sec; 'l4-conns-per-sec’: L4 Connections/sec; 'l7-conns-per-sec’: L7 Connections/sec; 'l7-trans-per-sec’: L7 Transactions/sec; 'ssl-conns-per-sec’: SSL Connections/sec; 'ip-nat-conns-per-sec’: IP NAT Connections/sec; 'total-new-conns-per-sec’: Total New Connections Established/sec; 'total-curr-conns’: Total Current Established Connections; 'l4-bandwidth’: L4 Bandwidth in bits/sec; 'l7-bandwidth’: L7 Bandwidth in bits/sec; 'serv-ssl-conns-per-sec’: Server SSL Connections/sec; 'fw-conns-per-sec’: FW Connections/sec; 'gifw-conns-per-sec’: GiFW Connections/sec; 'l7-proxy-conns-per-sec’: L7 Proxy Connections/sec; 'l7-proxy-trans-per-sec’: L7 Proxy Transactions/sec;