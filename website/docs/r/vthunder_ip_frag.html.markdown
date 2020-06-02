---
layout: "vthunder"
page_title: "vthunder: vthunder_ip_frag"
sidebar_current: "docs-vthunder-resource-ip-frag"
description: |-
	Provides details about vthunder ip frag resource for A10
---

# vthunder\_ip\_frag

`vthunder_ip_frag` Provides details about vthunder ip frag
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_ip_frag" "frag" {
  buff = 10000
  max_packets_per_reassembly = 2
  max_reassembly_sessions = 1
  timeout = 4
  sampling_enable  {
    counters1 = "all"
  }
  cpu_threshold  {
    high = 75
    low = 60
  }
}
```

## Argument Reference

* `buff` - Max buff used for fragmentation (Buffer Value(10000-3000000))
* `max_packets_per_reassembly` - Max number of fragmented packets allowed per reassembly(0 is unlimited) (default 0)
* `max_reassembly_sessions` - Max number of pending reassembly sessions allowed (default 100000)
* `timeout` - Fragmentation timeout (in milliseconds 4 - 16000 (default is 1000))
* `uuid` - uuid of the object
* `counters1` - ‘all’: all; ‘session-inserted’: Session Inserted; ‘session-expired’: Session Expired; ‘icmp-rcv’: ICMP Received; ‘icmpv6-rcv’: ICMPv6 Received; ‘udp-rcv’: UDP Received; ‘tcp-rcv’: TCP Received; ‘ipip-rcv’: IP-in-IP Received; ‘ipv6ip-rcv’: IPv6-in-IP Received; ‘other-rcv’: Other Received; ‘icmp-dropped’: ICMP Dropped; ‘icmpv6-dropped’: ICMPv6 Dropped; ‘udp-dropped’: UDP Dropped; ‘tcp-dropped’: TCP Dropped; ‘ipip-dropped’: IP-in-IP Dropped; ‘ipv6ip-dropped’: IPv6-in-IP Dropped; ‘other-dropped’: Other Dropped; ‘overlap-error’: Overlapping Fragment Dropped; ‘bad-ip-len’: Bad IP Length; ‘too-small’: Fragment Too Small Drop; ‘first-tcp-too-small’: First TCP Fragment Too Small Drop; ‘first-l4-too-small’: First L4 Fragment Too Small Drop; ‘total-sessions-exceeded’: Total Sessions Exceeded Drop; ‘no-session-memory’: Out of Session Memory; ‘fast-aging-set’: Fragmentation Fast Aging Set; ‘fast-aging-unset’: Fragmentation Fast Aging Unset; ‘fragment-queue-success’: Fragment Queue Success; ‘unaligned-len’: Payload Length Unaligned; ‘exceeded-len’: Payload Length Out of Bounds; ‘duplicate-first-frag’: Duplicate First Fragment; ‘duplicate-last-frag’: Duplicate Last Fragment; ‘total-fragments-exceeded’: Total Queued Fragments Exceeded; ‘fragment-queue-failure’: Fragment Queue Failure; ‘reassembly-success’: Fragment Reassembly Success; ‘max-len-exceeded’: Fragment Max Data Length Exceeded; ‘reassembly-failure’: Fragment Reassembly Failure; ‘policy-drop’: MTU Exceeded Policy Drop; ‘error-drop’: Fragment Processing Drop; ‘max-packets-exceeded’: Too Many Packets Per Reassembly Drop; ‘session-packets-exceeded’: Session Max Packets Exceeded; ‘frag-session-count’: Fragmentation Session Count; ‘high-cpu-threshold’: High CPU Threshold Reached; ‘low-cpu-threshold’: Low CPU Threshold Reached; ‘cpu-threshold-drop’: High CPU Drop; ‘ipd-entry-drop’: DDoS Protection Drop; ‘sctp-rcv’: SCTP Received; ‘sctp-dropped’: SCTP Dropped;
* `high` - When CPU usage reaches this value, it will stop processing fragments (default: 75%)
* `low` - When CPU usage remains under this value, it will resume processing fragments (default: 60%)

