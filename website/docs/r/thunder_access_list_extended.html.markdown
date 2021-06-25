---
layout: "thunder"
page_title: "thunder: thunder_access_list_extended"
sidebar_current: "docs-thunder-resource-access-list-extended"
description: |-
    Provides details about thunder access list extended resource for A10
---

# thunder\_access\_list\_extended

`thunder_access_list_extended` Provides details about thunder access list extended
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}




resource "thunder_access_list_extended" "resourceAccessListExtendedTest" {
	extd = 0
rules {   
	extd_seq_num =  0 
	extd_remark =  "string" 
	extd_action =  "string" 
	icmp =  0 
	tcp =  0 
	udp =  0 
	ip =  0 
	service_obj_group =  "string" 
	icmp_type =  0 
	any_type =  0 
	special_type =  "string" 
	any_code =  0 
	icmp_code =  0 
	special_code =  "string" 
	src_any =  0 
	src_host =  "string" 
	src_subnet =  "string" 
	src_mask =  "string" 
	src_object_group =  "string" 
	src_eq =  0 
	src_gt =  0 
	src_lt =  0 
	src_range =  0 
	src_port_end =  0 
	dst_any =  0 
	dst_host =  "string" 
	dst_subnet =  "string" 
	dst_mask =  "string" 
	dst_object_group =  "string" 
	dst_eq =  0 
	dst_gt =  0 
	dst_lt =  0 
	dst_range =  0 
	dst_port_end =  0 
	fragments =  0 
	vlan =  0 
	ethernet =  0 
	trunk =  0 
	dscp =  0 
	established =  0 
	acl_log =  0 
	transparent_session_only =  0 
	}
uuid = "string"
 
}

```

## Argument Reference

* `extended` - Configure Extended Access List
* `extd` - IP extended access list
* `extd-seq-num` - Sequence number
* `extd-remark` - Access list entry comment (Notes for this ACL)
* `extd-action` - 'deny': Deny; 'permit': Permit; 'l3-vlan-fwd-disable': Disable L3 forwarding between VLANs;
* `icmp` - Internet Control Message Protocol
* `tcp` - protocol TCP
* `udp` - protocol UDP
* `ip` - Any Internet Protocol
* `service-obj-group` - Service object group (Source object group name)
* `icmp-type` - ICMP type number
* `any-type` - Any ICMP type
* `special-type` - 'echo-reply': Type 0, echo reply; 'echo-request': Type 8, echo request; 'info-reply': Type 16, information reply; 'info-request': Type 15, information request; 'mask-reply': Type 18, address mask reply; 'mask-request': Type 17, address mask request; 'parameter-problem': Type 12, parameter problem; 'redirect': Type 5, redirect message; 'source-quench': Type 4, source quench; 'time-exceeded': Type 11, time exceeded; 'timestamp': Type 13, timestamp; 'timestamp-reply': Type 14, timestamp reply; 'dest-unreachable': Type 3, destination unreachable;
* `any-code` - Any ICMP code
* `icmp-code` - ICMP code number
* `special-code` - 'frag-required': Code 4, fragmentation required; 'host-unreachable': Code 1, destination host unreachable; 'network-unreachable': Code 0, destination network unreachable; 'port-unreachable': Code 3, destination port unreachable; 'proto-unreachable': Code 2, destination protocol unreachable; 'route-failed': Code 5, source route failed;
* `src-any` - Any source host
* `src-host` - A single source host (Host address)
* `src-subnet` - Source Address
* `src-mask` - Source Mask 0=apply 255=ignore
* `src-object-group` - Network object group (Source network object group name)
* `src-eq` - Match only packets on a given source port (port number)
* `src-gt` - Match only packets with a greater port number
* `src-lt` - Match only packets with a lower port number
* `src-range` - match only packets in the range of port numbers (Starting Port Number)
* `src-port-end` - Ending Port Number
* `dst-any` - Any destination host
* `dst-host` - A single destination host (Host address)
* `dst-subnet` - Destination Address
* `dst-mask` - Destination Mask 0=apply 255=ignore
* `dst-object-group` - Destination network object group name
* `dst-eq` - Match only packets on a given destination port (port number)
* `dst-gt` - Match only packets with a greater port number
* `dst-lt` - Match only packets with a lesser port number
* `dst-range` - Match only packets in the range of port numbers (Starting Destination Port Number)
* `dst-port-end` - Edning Destination Port Number
* `fragments` - IP fragments
* `vlan` - VLAN ID
* `ethernet` - Ethernet interface (Port number)
* `trunk` - Ethernet trunk (trunk number)
* `dscp` - DSCP
* `established` - TCP established
* `acl-log` - Log matches against this entry
* `transparent-session-only` - Only log transparent sessions
* `uuid` - uuid of the object

