---
layout: "thunder"
page_title: "thunder: thunder_rule_set"
sidebar_current: "docs-thunder-resource-rule-set"
description: |-
    Provides details about thunder rule set resource for A10
---

# thunder\_rule\_set

`thunder_rule_set` Provides details about thunder rule set
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_rule_set" "resourceRuleSetTest" {
	name = "string"
session_statistic = "string"
remark = "string"
uuid = "string"
user_tag = "string"
sampling-enable {   
	counters1 =  "string" 
	}
packet_capture_template = "string"
rule-list {   
	name =  "string" 
	remark =  "string" 
	status =  "string" 
	action =  "string" 
	log =  0 
	reset_lid =  0 
	listen_on_port =  0 
	policy =  "string" 
	vpn_ipsec_name =  "string" 
	forward_listen_on_port =  0 
	lid =  0 
	listen_on_port_lid =  0 
	fw_log =  0 
	fwlog =  0 
	cgnv6_log =  0 
	forward_log =  0 
	lidlog =  0 
	reset_lidlog =  0 
	listen_on_port_lidlog =  0 
	cgnv6_policy =  "string" 
	cgnv6_fixed_nat_log =  0 
	cgnv6_lsn_lid =  0 
	cgnv6_lsn_log =  0 
	ip_version =  "string" 
	gtp_template =  "string" 
	src_class_list =  "string" 
	src_geoloc_name =  "string" 
	src_geoloc_list =  "string" 
	src_geoloc_list_shared =  0 
	src_ipv4_any =  "string" 
	src_ipv6_any =  "string" 
source-list {   
	src_ip_subnet =  "string" 
	src_ipv6_subnet =  "string" 
	src_obj_network =  "string" 
	src_obj_grp_network =  "string" 
	src_slb_server =  "string" 
	}
	src_zone =  "string" 
	src_zone_any =  "string" 
	src_threat_list =  "string" 
	dst_class_list =  "string" 
	dst_geoloc_name =  "string" 
	dst_geoloc_list =  "string" 
	dst_geoloc_list_shared =  0 
	dst_ipv4_any =  "string" 
	dst_ipv6_any =  "string" 
dest-list {   
	dst_ip_subnet =  "string" 
	dst_ipv6_subnet =  "string" 
	dst_obj_network =  "string" 
	dst_obj_grp_network =  "string" 
	dst_slb_server =  "string" 
	dst_slb_vserver =  "string" 
	}
	dst_domain_list =  "string" 
	dst_zone =  "string" 
	dst_zone_any =  "string" 
	dst_threat_list =  "string" 
	service_any =  "string" 
service-list {   
	protocols =  "string" 
	proto_id =  0 
	obj_grp_service =  "string" 
	icmp =  0 
	icmpv6 =  0 
	icmp_type =  0 
	special_type =  "string" 
	icmp_code =  0 
	special_code =  "string" 
	icmpv6_type =  0 
	special_v6_type =  "string" 
	icmpv6_code =  0 
	special_v6_code =  "string" 
	eq_src_port =  0 
	gt_src_port =  0 
	lt_src_port =  0 
	range_src_port =  0 
	port_num_end_src =  0 
	eq_dst_port =  0 
	gt_dst_port =  0 
	lt_dst_port =  0 
	range_dst_port =  0 
	port_num_end_dst =  0 
	sctp_template =  "string" 
	alg =  "string" 
	}
	idle_timeout =  0 
dscp-list {   
	dscp_value =  "string" 
	dscp_range_start =  0 
	dscp_range_end =  0 
	}
	application_any =  "string" 
app-list {   
	obj_grp_application =  "string" 
	protocol =  "string" 
	protocol_tag =  "string" 
	}
	track_application =  0 
	uuid =  "string" 
	user_tag =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	packet_capture_template =  "string" 
action_group {  
 	type =  "string" 
	permit_log =  0 
	reset_log =  0 
	deny_log =  0 
	listen_on_port =  0 
	forward =  0 
	ipsec =  0 
	vpn_ipsec_name =  "string" 
	cgnv6 =  0 
	cgnv6_policy =  "string" 
	cgnv6_lsn_lid =  0 
	permit_limit_policy =  0 
	permit_respond_to_user_mac =  0 
	reset_respond_to_user_mac =  0 
	set_dscp =  0 
	dscp_value =  "string" 
	dscp_number =  0 
	uuid =  "string" 
	}
move_rule {  
 	location =  "string" 
	target_rule =  "string" 
	}
	}
rules_by_zone {  
 	uuid =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
application {  
 	uuid =  "string" 
	}
track_app_rule_list {  
 	uuid =  "string" 
	}
app {  
 	uuid =  "string" 
	}
tag {  
 	uuid =  "string" 
	}
 
}

```

## Argument Reference

* `rule-set` - Configure Security policy Rule Set
* `name` - Rule name
* `session-statistic` - 'enable': Enable session based statistic (Default); 'disable': Disable session based statistic;
* `remark` - Rule entry comment (Notes for this rule)
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `counters1` - 'all': all; 'dummy': Entry for a10countergen;
* `packet-capture-template` - Name of the packet capture template to be bind with this object
* `status` - 'enable': Enable rule; 'disable': Disable rule;
* `action` - 'permit': permit; 'deny': deny; 'reset': reset;
* `log` - Enable logging
* `reset-lid` - Apply a Template LID
* `listen-on-port` - Listen on port
* `policy` - 'cgnv6': Apply CGNv6 policy; 'forward': Forward packet; 'ipsec': Apply IPsec encapsulation;
* `vpn-ipsec-name` - VPN IPsec name
* `forward-listen-on-port` - Listen on port
* `lid` - Apply a Template LID
* `listen-on-port-lid` - Apply a Template LID
* `fw-log` - Enable logging
* `fwlog` - Enable logging
* `cgnv6-log` - Enable logging
* `forward-log` - Enable logging
* `lidlog` - Enable logging
* `reset-lidlog` - Enable logging
* `listen-on-port-lidlog` - Enable logging
* `cgnv6-policy` - 'lsn-lid': Apply specified CGNv6 LSN LID; 'fixed-nat': Apply CGNv6 Fixed NAT;
* `cgnv6-fixed-nat-log` - Enable logging
* `cgnv6-lsn-lid` - LSN LID
* `cgnv6-lsn-log` - Enable logging
* `ip-version` - 'v4': IPv4 rule; 'v6': IPv6 rule;
* `gtp-template` - Configure GTP Policy Template (GTP Template Policy Name)
* `src-class-list` - Match source IP against class-list
* `src-geoloc-name` - Single geolocation name
* `src-geoloc-list` - Geolocation name list
* `src-geoloc-list-shared` - Use Geolocation list from shared partition
* `src-ipv4-any` - 'any': Any IPv4 address;
* `src-ipv6-any` - 'any': Any IPv6 address;
* `src-ip-subnet` - IPv4 IP Address
* `src-ipv6-subnet` - IPv6 IP Address
* `src-obj-network` - Network object
* `src-obj-grp-network` - Network object group
* `src-slb-server` - SLB Real server name
* `src-zone` - Zone name
* `src-zone-any` - 'any': any;
* `src-threat-list` - Bind threat-list for source IP based filtering
* `dst-class-list` - Match destination IP against class-list
* `dst-geoloc-name` - Single geolocation name
* `dst-geoloc-list` - Geolocation name list
* `dst-geoloc-list-shared` - Use Geolocation list from shared partition
* `dst-ipv4-any` - 'any': Any IPv4 address;
* `dst-ipv6-any` - 'any': Any IPv6 address;
* `dst-ip-subnet` - IPv4 IP Address
* `dst-ipv6-subnet` - IPv6 IP Address
* `dst-obj-network` - Network object
* `dst-obj-grp-network` - Network object group
* `dst-slb-server` - SLB Real server name
* `dst-slb-vserver` - SLB Virtual server name
* `dst-domain-list` - Match destination IP against domain-list
* `dst-zone` - Zone name
* `dst-zone-any` - 'any': any;
* `dst-threat-list` - Bind threat-list for destination IP based filtering
* `service-any` - 'any': any;
* `protocols` - 'tcp': tcp; 'udp': udp; 'sctp': sctp;
* `proto-id` - Protocol ID
* `obj-grp-service` - service object group
* `icmp` - ICMP
* `icmpv6` - ICMPv6
* `icmp-type` - ICMP type number
* `special-type` - 'any-type': Any ICMP type; 'echo-reply': Type 0, echo reply; 'echo-request': Type 8, echo request; 'info-reply': Type 16, information reply; 'info-request': Type 15, information request; 'mask-reply': Type 18, address mask reply; 'mask-request': Type 17, address mask request; 'parameter-problem': Type 12, parameter problem; 'redirect': Type 5, redirect message; 'source-quench': Type 4, source quench; 'time-exceeded': Type 11, time exceeded; 'timestamp': Type 13, timestamp; 'timestamp-reply': Type 14, timestamp reply; 'dest-unreachable': Type 3, destination unreachable;
* `icmp-code` - ICMP code number
* `special-code` - 'any-code': Any ICMP code; 'frag-required': Code 4, fragmentation required; 'host-unreachable': Code 1, destination host unreachable; 'network-unreachable': Code 0, destination network unreachable; 'port-unreachable': Code 3, destination port unreachable; 'proto-unreachable': Code 2, destination protocol unreachable; 'route-failed': Code 5, source route failed;
* `icmpv6-type` - ICMPv6 type number
* `special-v6-type` - 'any-type': Any ICMPv6 type; 'dest-unreachable': Type 1, destination unreachable; 'echo-reply': Type 129, echo reply; 'echo-request': Type 128, echo request; 'packet-too-big': Type 2, packet too big; 'param-prob': Type 4, parameter problem; 'time-exceeded': Type 3, time exceeded;
* `icmpv6-code` - ICMPv6 code number
* `special-v6-code` - 'any-code': Any ICMPv6 code; 'addr-unreachable': Code 3, address unreachable; 'admin-prohibited': Code 1, admin prohibited; 'no-route': Code 0, no route to destination; 'not-neighbour': Code 2, not neighbor; 'port-unreachable': Code 4, destination port unreachable;
* `eq-src-port` - Equal to the port number
* `gt-src-port` - Greater than the port number
* `lt-src-port` - Lower than the port number
* `range-src-port` - Port range (Starting Port Number)
* `port-num-end-src` - Ending Port Number
* `eq-dst-port` - Equal to the port number
* `gt-dst-port` - Greater than the port number
* `lt-dst-port` - Lower than the port number
* `range-dst-port` - Port range (Starting Port Number)
* `port-num-end-dst` - Ending Port Number
* `sctp-template` - SCTP Template
* `alg` - 'FTP': FTP; 'TFTP': TFTP; 'SIP': SIP; 'DNS': DNS; 'PPTP': PPTP; 'RTSP': RTSP; 'ESP': ESP;
* `idle-timeout` - TCP/UDP idle-timeout
* `dscp-value` - 'default': Default dscp (000000); 'af11': AF11 (001010); 'af12': AF12 (001100); 'af13': AF13 (001110); 'af21': AF21 (010010); 'af22': AF22 (010100); 'af23': AF23 (010110); 'af31': AF31 (011010); 'af32': AF32 (011100); 'af33': AF33 (011110); 'af41': AF41 (100010); 'af42': AF42 (100100); 'af43': AF43 (100110); 'cs1': CS1 (001000); 'cs2': CS2 (010000); 'cs3': CS3 (011000); 'cs4': CS4 (100000); 'cs5': CS5 (101000); 'cs6': CS6 (110000); 'cs7': CS7 (111000); 'ef': EF (101110);
* `dscp-range-start` - Start DSCP Number
* `dscp-range-end` - Ending DSCP Number
* `application-any` - 'any': any;
* `obj-grp-application` - Application object group
* `protocol` - Specify application(s)
* `protocol-tag` - 'aaa': Protocol/application used for AAA (Authentification, Authorization and Accounting) purposes.; 'adult-content': Adult content protocol/application.; 'advertising': Advertising networks and applications.; 'aetls': Application known to enforce HSTS and thus use of TLS.; 'analytics-and-statistics': User analytics and statistics protocol/application.; 'anonymizers-and-proxies': Traffic-anonymization protocol/application.; 'audio-chat': Protocol/application used for Audio Chat.; 'basic': Covers all protocols required for basic classification, including most networking protocols as well as standard protocols like HTTP.; 'blog': Blogging platform protocol/application.; 'cdn': Protocol/application used for Content-Delivery Networks.; 'certification-authority': Certification Authority for SSL/TLS certificate.; 'chat': Protocol/application used for Text Chat.; 'classified-ads': Protocol/application used for Classified Advertisements.; 'cloud-based-services': SaaS and/or PaaS cloud based services.; 'crowdfunding': Service for funding a project or venture by raising small amounts of money from a large number of people, typically via the Internet.; 'cryptocurrency': Services for mining cryptocurrencies, for example a Crypto Web Browser (an application that mines crypto currency in the background while its user browses the web).; 'database': Database-specific protocols.; 'disposable-email': Service offering Disposable Email Accounts (DEA). DEA is a technique to share temporary email address between many users.; 'ebook-reader': Services for e-book readers, i.e. connected devices that display electronic books (typically using e-ink displays to reduce glare and eye strain).; 'education': Protocols offering education services and online courses.; 'email': Native email protocol.; 'enterprise': Protocol/application used in an enterprise network.; 'file-management': Protocol/application designed specifically for file management and exchange. This can include bona fide network protocols (like SMB) as well as web/cloud services (like Dropbox).; 'file-transfer': Protocol that offers file transferring as a secondary feature. This typically includes IM, WebMail, and other protocols that allow file transfers in addition to their principal function.; 'forum': Online forum protocol/application.; 'gaming': Protocol/application used by games.; 'healthcare': Protocols offering medical services, i.e protocols used in medical environment.; 'instant-messaging-and-multimedia-conferencing': Protocol/application used for Instant Messaging or Multi-Conferencing.; 'internet-of-things': Internet Of Things protocol/application.; 'map-service': Digital Maps service (web site and their related API).; 'mobile': Mobile-specific protocol/application.; 'multimedia-streaming': Protocol/application used for multimedia streaming.; 'networking': Protocol used for (inter) networking purpose.; 'news-portal': Protocol/application used for News Portals.; 'payment-service': Application offering online services for accepting electronic payments by a variety of payment methods (credit card, bank-based payments such as direct debit, bank transfer, etc).; 'peer-to-peer': Protocol/application used for Peer-to-peer purposes.; 'remote-access': Protocol/application used for remote access.; 'scada': SCADA (Supervisory control and data acquisition) protocols, all generations.; 'social-networks': Social networking application.; 'software-update': Auto-update protocol.; 'speedtest': Speedtest application allowing to access quality of Internet connection (upload, download, latency, etc).; 'standards-based': Protocol issued from standardized bodies such as IETF, ITU, IEEE, ETSI, OIF.; 'transportation': Transportation services, for example smartphone applications that allow users to hail a taxi.; 'video-chat': Protocol/application used for Video Chat.; 'voip': Application used for Voice-Over-IP.; 'vpn-tunnels': Protocol/application used for VPN or tunneling purposes.; 'web': Application based on HTTP/HTTPS.; 'web-e-commerce': Protocol/application used for E-commerce websites.; 'web-search-engines': Protocol/application used for Web search portals.; 'web-websites': Protocol/application used for Company Websites.; 'webmails': Web-based e-mail application.; 'web-ext-adult': Web Extension Adult; 'web-ext-auctions': Web Extension Auctions; 'web-ext-blogs': Web Extension Blogs; 'web-ext-business-and-economy': Web Extension Business and Economy; 'web-ext-cdns': Web Extension CDNs; 'web-ext-collaboration': Web Extension Collaboration; 'web-ext-computer-and-internet-info': Web Extension Computer and Internet Info; 'web-ext-computer-and-internet-security': Web Extension Computer and Internet Security; 'web-ext-dating': Web Extension Dating; 'web-ext-educational-institutions': Web Extension Educational Institutions; 'web-ext-entertainment-and-arts': Web Extension Entertainment and Arts; 'web-ext-fashion-and-beauty': Web Extension Fashion and Beauty; 'web-ext-file-share': Web Extension File Share; 'web-ext-financial-services': Web Extension Financial Services; 'web-ext-gambling': Web Extension Gambling; 'web-ext-games': Web Extension Games; 'web-ext-government': Web Extension Government; 'web-ext-health-and-medicine': Web Extension Health and Medicine; 'web-ext-individual-stock-advice-and-tools': Web Extension Individual Stock Advice and Tools; 'web-ext-internet-portals': Web Extension Internet Portals; 'web-ext-job-search': Web Extension Job Search; 'web-ext-local-information': Web Extension Local Information; 'web-ext-malware': Web Extension Malware; 'web-ext-motor-vehicles': Web Extension Motor Vehicles; 'web-ext-music': Web Extension Music; 'web-ext-news': Web Extension News; 'web-ext-p2p': Web Extension P2P; 'web-ext-parked-sites': Web Extension Parked Sites; 'web-ext-proxy-avoid-and-anonymizers': Web Extension Proxy Avoid and Anonymizers; 'web-ext-real-estate': Web Extension Real Estate; 'web-ext-reference-and-research': Web Extension Reference and Research; 'web-ext-search-engines': Web Extension Search Engines; 'web-ext-shopping': Web Extension Shopping; 'web-ext-social-network': Web Extension Social Network; 'web-ext-society': Web Extension Society; 'web-ext-software': Web Extension Software; 'web-ext-sports': Web Extension Sports; 'web-ext-streaming-media': Web Extension Streaming Media; 'web-ext-training-and-tools': Web Extension Training and Tools; 'web-ext-translation': Web Extension Translation; 'web-ext-travel': Web Extension Travel; 'web-ext-web-advertisements': Web Extension Web Advertisements; 'web-ext-web-based-email': Web Extension Web based Email; 'web-ext-web-hosting': Web Extension Web Hosting; 'web-ext-web-service': Web Extension Web Service;
* `track-application` - Enable application statistic
* `type` - 'permit': permit; 'deny': deny; 'reset': reset;
* `permit-log` - Enable logging
* `reset-log` - Enable logging
* `deny-log` - Enable logging
* `forward` - Forward packet
* `ipsec` - Apply IPsec encapsulation
* `cgnv6` - Apply CGNv6 policy
* `permit-limit-policy` - Limit policy Template
* `permit-respond-to-user-mac` - Use the user's source MAC for the next hop rather than the routing table (default:off)
* `reset-respond-to-user-mac` - Use the user's source MAC for the next hop rather than the routing table (default:off)
* `set-dscp` - DSCP setting
* `dscp-number` - DSCP Number
* `location` - 'top': top; 'before': before; 'after': after; 'bottom': bottom;

