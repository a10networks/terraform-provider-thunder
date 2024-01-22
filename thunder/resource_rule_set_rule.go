package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_rule_set_rule`: Configure rule-set rule\n\n__PLACEHOLDER__",
		CreateContext: resourceRuleSetRuleCreate,
		UpdateContext: resourceRuleSetRuleUpdate,
		ReadContext:   resourceRuleSetRuleRead,
		DeleteContext: resourceRuleSetRuleDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'permit': permit; 'deny': deny; 'reset': reset;",
			},
			"action_group": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Optional: true, Description: "'permit': permit; 'deny': deny; 'reset': reset;",
						},
						"permit_log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
						},
						"reset_log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
						},
						"deny_log": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
						},
						"logging_template_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"permit_log_template_type": {
										Type: schema.TypeString, Optional: true, Description: "'fw-logging-template': Logging with specified fw template; 'cgnv6-logging-template': Logging with specified cgnv6 template; 'netflow-monitor': Logging with specified netflow/ipfix monitor;",
									},
									"permit_fw_log": {
										Type: schema.TypeString, Optional: true, Description: "Logging template name",
									},
									"permit_cgnv6_log": {
										Type: schema.TypeString, Optional: true, Description: "Logging template name",
									},
									"permit_netflow_log": {
										Type: schema.TypeString, Optional: true, Description: "Name of netflow monitor",
									},
								},
							},
						},
						"reset_log_template_type": {
							Type: schema.TypeString, Optional: true, Description: "'fw-logging-template': Logging with specified fw template;",
						},
						"reset_fw_log": {
							Type: schema.TypeString, Optional: true, Description: "Logging template name",
						},
						"deny_log_template_type": {
							Type: schema.TypeString, Optional: true, Description: "'fw-logging-template': Logging with specified fw template;",
						},
						"deny_fw_log": {
							Type: schema.TypeString, Optional: true, Description: "Logging template name",
						},
						"listen_on_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Listen on port",
						},
						"forward": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward packet",
						},
						"ipsec": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply IPsec encapsulation",
						},
						"ipsec_group": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply IPsec Group encapsulation",
						},
						"vpn_ipsec_name": {
							Type: schema.TypeString, Optional: true, Description: "VPN IPsec name",
						},
						"vpn_ipsec_group_name": {
							Type: schema.TypeString, Optional: true, Description: "VPN IPsec Group name",
						},
						"cgnv6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply CGNv6 policy",
						},
						"cgnv6_policy": {
							Type: schema.TypeString, Optional: true, Description: "'lsn-lid': Apply specified CGNv6 LSN LID; 'fixed-nat': Apply CGNv6 Fixed NAT; 'ds-lite': Apply CGNv6 DS-Lite;",
						},
						"cgnv6_lsn_lid": {
							Type: schema.TypeInt, Optional: true, Description: "LSN LID",
						},
						"cgnv6_ds_lite": {
							Type: schema.TypeString, Optional: true, Description: "'lsn-lid': Apply specified CGNv6 LSN LID;",
						},
						"cgnv6_ds_lite_lsn_lid": {
							Type: schema.TypeInt, Optional: true, Description: "LSN LID",
						},
						"inspect_payload": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DS-Lite tunnel inspection",
						},
						"permit_limit_policy": {
							Type: schema.TypeInt, Optional: true, Description: "Limit policy Template",
						},
						"permit_respond_to_user_mac": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the user's source MAC for the next hop rather than the routing table (default:off)",
						},
						"reset_respond_to_user_mac": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the user's source MAC for the next hop rather than the routing table (default:off)",
						},
						"set_dscp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DSCP setting",
						},
						"dscp_value": {
							Type: schema.TypeString, Optional: true, Description: "'default': Default dscp (000000); 'af11': AF11 (001010); 'af12': AF12 (001100); 'af13': AF13 (001110); 'af21': AF21 (010010); 'af22': AF22 (010100); 'af23': AF23 (010110); 'af31': AF31 (011010); 'af32': AF32 (011100); 'af33': AF33 (011110); 'af41': AF41 (100010); 'af42': AF42 (100100); 'af43': AF43 (100110); 'cs1': CS1 (001000); 'cs2': CS2 (010000); 'cs3': CS3 (011000); 'cs4': CS4 (100000); 'cs5': CS5 (101000); 'cs6': CS6 (110000); 'cs7': CS7 (111000); 'ef': EF (101110);",
						},
						"dscp_number": {
							Type: schema.TypeInt, Optional: true, Description: "DSCP Number",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"app_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"obj_grp_application": {
							Type: schema.TypeString, Optional: true, Description: "Application object group",
						},
						"protocol": {
							Type: schema.TypeString, Optional: true, Description: "Specify application(s)",
						},
						"protocol_tag": {
							Type: schema.TypeString, Optional: true, Description: "'aaa': Protocol/application used for AAA (Authentification, Authorization and Accounting) purposes.; 'adult-content': Adult content protocol/application.; 'advertising': Advertising networks and applications.; 'application-enforcing-tls': Application known to enforce HSTS and thus use of TLS.; 'analytics-and-statistics': User analytics and statistics protocol/application.; 'anonymizers-and-proxies': Traffic-anonymization protocol/application.; 'audio-chat': Protocol/application used for Audio Chat.; 'basic': Covers all protocols required for basic classification, including most networking protocols as well as standard protocols like HTTP.; 'blog': Blogging platform protocol/application.; 'cdn': Protocol/application used for Content-Delivery Networks.; 'certification-authority': Certification Authority for SSL/TLS certificate.; 'chat': Protocol/application used for Text Chat.; 'classified-ads': Protocol/application used for Classified Advertisements.; 'cloud-based-services': SaaS and/or PaaS cloud based services.; 'crowdfunding': Service for funding a project or venture by raising small amounts of money from a large number of people, typically via the Internet.; 'cryptocurrency': Services for mining cryptocurrencies, for example a Crypto Web Browser (an application that mines crypto currency in the background while its user browses the web).; 'database': Database-specific protocols.; 'disposable-email': Service offering Disposable Email Accounts (DEA). DEA is a technique to share temporary email address between many users.; 'ebook-reader': Services for e-book readers, i.e. connected devices that display electronic books (typically using e-ink displays to reduce glare and eye strain).; 'education': Protocols offering education services and online courses.; 'email': Native email protocol.; 'enterprise': Protocol/application used in an enterprise network.; 'file-management': Protocol/application designed specifically for file management and exchange. This can include bona fide network protocols (like SMB) as well as web/cloud services (like Dropbox).; 'file-transfer': Protocol that offers file transferring as a secondary feature. This typically includes IM, WebMail, and other protocols that allow file transfers in addition to their principal function.; 'forum': Online forum protocol/application.; 'gaming': Protocol/application used by games.; 'healthcare': Protocols offering medical services, i.e protocols used in medical environment.; 'instant-messaging-and-multimedia-conferencing': Protocol/application used for Instant Messaging or Multi-Conferencing.; 'internet-of-things': Internet Of Things protocol/application.; 'map-service': Digital Maps service (web site and their related API).; 'mobile': Mobile-specific protocol/application.; 'multimedia-streaming': Protocol/application used for multimedia streaming.; 'networking': Protocol used for (inter) networking purpose.; 'news-portal': Protocol/application used for News Portals.; 'payment-service': Application offering online services for accepting electronic payments by a variety of payment methods (credit card, bank-based payments such as direct debit, bank transfer, etc).; 'peer-to-peer': Protocol/application used for Peer-to-peer purposes.; 'remote-access': Protocol/application used for remote access.; 'scada': SCADA (Supervisory control and data acquisition) protocols, all generations.; 'social-networks': Social networking application.; 'software-update': Auto-update protocol.; 'speedtest': Speedtest application allowing to access quality of Internet connection (upload, download, latency, etc).; 'standards-based': Protocol issued from standardized bodies such as IETF, ITU, IEEE, ETSI, OIF.; 'transportation': Transportation services, for example smartphone applications that allow users to hail a taxi.; 'video-chat': Protocol/application used for Video Chat.; 'voip': Application used for Voice-Over-IP.; 'vpn-tunnels': Protocol/application used for VPN or tunneling purposes.; 'web': Application based on HTTP/HTTPS.; 'web-e-commerce': Protocol/application used for E-commerce websites.; 'web-search-engines': Protocol/application used for Web search portals.; 'web-websites': Protocol/application used for Company Websites.; 'webmails': Web-based e-mail application.; 'web-ext-adult': Web Extension Adult; 'web-ext-auctions': Web Extension Auctions; 'web-ext-blogs': Web Extension Blogs; 'web-ext-business-and-economy': Web Extension Business and Economy; 'web-ext-cdns': Web Extension CDNs; 'web-ext-collaboration': Web Extension Collaboration; 'web-ext-computer-and-internet-info': Web Extension Computer and Internet Info; 'web-ext-computer-and-internet-security': Web Extension Computer and Internet Security; 'web-ext-dating': Web Extension Dating; 'web-ext-educational-institutions': Web Extension Educational Institutions; 'web-ext-entertainment-and-arts': Web Extension Entertainment and Arts; 'web-ext-fashion-and-beauty': Web Extension Fashion and Beauty; 'web-ext-file-share': Web Extension File Share; 'web-ext-financial-services': Web Extension Financial Services; 'web-ext-gambling': Web Extension Gambling; 'web-ext-games': Web Extension Games; 'web-ext-government': Web Extension Government; 'web-ext-health-and-medicine': Web Extension Health and Medicine; 'web-ext-individual-stock-advice-and-tools': Web Extension Individual Stock Advice and Tools; 'web-ext-internet-portals': Web Extension Internet Portals; 'web-ext-job-search': Web Extension Job Search; 'web-ext-local-information': Web Extension Local Information; 'web-ext-malware': Web Extension Malware; 'web-ext-motor-vehicles': Web Extension Motor Vehicles; 'web-ext-music': Web Extension Music; 'web-ext-news': Web Extension News; 'web-ext-p2p': Web Extension P2P; 'web-ext-parked-sites': Web Extension Parked Sites; 'web-ext-proxy-avoid-and-anonymizers': Web Extension Proxy Avoid and Anonymizers; 'web-ext-real-estate': Web Extension Real Estate; 'web-ext-reference-and-research': Web Extension Reference and Research; 'web-ext-search-engines': Web Extension Search Engines; 'web-ext-shopping': Web Extension Shopping; 'web-ext-social-network': Web Extension Social Network; 'web-ext-society': Web Extension Society; 'web-ext-software': Web Extension Software; 'web-ext-sports': Web Extension Sports; 'web-ext-streaming-media': Web Extension Streaming Media; 'web-ext-training-and-tools': Web Extension Training and Tools; 'web-ext-translation': Web Extension Translation; 'web-ext-travel': Web Extension Travel; 'web-ext-web-advertisements': Web Extension Web Advertisements; 'web-ext-web-based-email': Web Extension Web based Email; 'web-ext-web-hosting': Web Extension Web Hosting; 'web-ext-web-service': Web Extension Web Service;",
						},
					},
				},
			},
			"application_any": {
				Type: schema.TypeString, Optional: true, Default: "any", Description: "'any': any;",
			},
			"cgnv6_ds_lite": {
				Type: schema.TypeString, Optional: true, Description: "'lsn-lid': Apply specified CGNv6 LSN LID;",
			},
			"cgnv6_ds_lite_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"cgnv6_ds_lite_lsn_lid": {
				Type: schema.TypeInt, Optional: true, Description: "LSN LID",
			},
			"cgnv6_fixed_nat_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"cgnv6_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"cgnv6_lsn_lid": {
				Type: schema.TypeInt, Optional: true, Description: "LSN LID",
			},
			"cgnv6_lsn_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"cgnv6_policy": {
				Type: schema.TypeString, Optional: true, Description: "'lsn-lid': Apply specified CGNv6 LSN LID; 'fixed-nat': Apply CGNv6 Fixed NAT; 'ds-lite': Apply CGNv6 DS-Lite;",
			},
			"dest_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_ip_subnet": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 IP Address",
						},
						"dst_ipv6_subnet": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 IP Address",
						},
						"dst_obj_network": {
							Type: schema.TypeString, Optional: true, Description: "Network object",
						},
						"dst_obj_grp_network": {
							Type: schema.TypeString, Optional: true, Description: "Network object group",
						},
						"dst_slb_server": {
							Type: schema.TypeString, Optional: true, Description: "SLB Real server name",
						},
						"dst_slb_vserver": {
							Type: schema.TypeString, Optional: true, Description: "SLB Virtual server name",
						},
					},
				},
			},
			"dscp_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp_value": {
							Type: schema.TypeString, Optional: true, Description: "'default': Default dscp (000000); 'af11': AF11 (001010); 'af12': AF12 (001100); 'af13': AF13 (001110); 'af21': AF21 (010010); 'af22': AF22 (010100); 'af23': AF23 (010110); 'af31': AF31 (011010); 'af32': AF32 (011100); 'af33': AF33 (011110); 'af41': AF41 (100010); 'af42': AF42 (100100); 'af43': AF43 (100110); 'cs1': CS1 (001000); 'cs2': CS2 (010000); 'cs3': CS3 (011000); 'cs4': CS4 (100000); 'cs5': CS5 (101000); 'cs6': CS6 (110000); 'cs7': CS7 (111000); 'ef': EF (101110);",
						},
						"dscp_range_start": {
							Type: schema.TypeInt, Optional: true, Description: "Start DSCP Number",
						},
						"dscp_range_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ending DSCP Number",
						},
					},
				},
			},
			"dst_class_list": {
				Type: schema.TypeString, Optional: true, Description: "Match destination IP against class-list",
			},
			"dst_domain_list": {
				Type: schema.TypeString, Optional: true, Description: "Match destination IP against domain-list",
			},
			"dst_geoloc_list": {
				Type: schema.TypeString, Optional: true, Description: "Geolocation name list",
			},
			"dst_geoloc_list_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Geolocation list from shared partition",
			},
			"dst_geoloc_name": {
				Type: schema.TypeString, Optional: true, Description: "Single geolocation name",
			},
			"dst_ipv4_any": {
				Type: schema.TypeString, Optional: true, Default: "any", Description: "'any': Any IPv4 address;",
			},
			"dst_ipv6_any": {
				Type: schema.TypeString, Optional: true, Default: "any", Description: "'any': Any IPv6 address;",
			},
			"dst_threat_list": {
				Type: schema.TypeString, Optional: true, Description: "Bind threat-list for destination IP based filtering",
			},
			"dst_zone": {
				Type: schema.TypeString, Optional: true, Description: "Zone name",
			},
			"dst_zone_any": {
				Type: schema.TypeString, Optional: true, Default: "any", Description: "'any': any;",
			},
			"forward_listen_on_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Listen on port",
			},
			"forward_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"fw_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"fwlog": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"gtp_template": {
				Type: schema.TypeString, Optional: true, Description: "Configure GTP Policy Template (GTP Template Policy Name)",
			},
			"idle_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "TCP/UDP idle-timeout",
			},
			"inspect_payload": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable DS-Lite tunnel inspection",
			},
			"ip_version": {
				Type: schema.TypeString, Optional: true, Default: "v4", Description: "'v4': IPv4 rule; 'v6': IPv6 rule;",
			},
			"lid": {
				Type: schema.TypeInt, Optional: true, Description: "Apply a Template LID",
			},
			"lidlog": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"listen_on_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Listen on port",
			},
			"listen_on_port_lid": {
				Type: schema.TypeInt, Optional: true, Description: "Apply a Template LID",
			},
			"listen_on_port_lidlog": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"move_rule": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"location": {
							Type: schema.TypeString, Optional: true, Default: "bottom", Description: "'top': top; 'before': before; 'after': after; 'bottom': bottom;",
						},
						"target_rule": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Rule name",
			},
			"policy": {
				Type: schema.TypeString, Optional: true, Description: "'cgnv6': Apply CGNv6 policy; 'forward': Forward packet; 'ipsec': Apply IPsec encapsulation; 'ipsec-group': Apply IPsec encapsulation from a group;",
			},
			"remark": {
				Type: schema.TypeString, Optional: true, Description: "Rule entry comment (Notes for this rule)",
			},
			"reset_lid": {
				Type: schema.TypeInt, Optional: true, Description: "Apply a Template LID",
			},
			"reset_lidlog": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'hit-count': Hit counts; 'permit-bytes': Permitted bytes counter; 'deny-bytes': Denied bytes counter; 'reset-bytes': Reset bytes counter; 'permit-packets': Permitted packets counter; 'deny-packets': Denied packets counter; 'reset-packets': Reset packets counter; 'active-session-tcp': Active TCP session counter; 'active-session-udp': Active UDP session counter; 'active-session-icmp': Active ICMP session counter; 'active-session-other': Active other protocol session counter; 'session-tcp': TCP session counter; 'session-udp': UDP session counter; 'session-icmp': ICMP session counter; 'session-other': Other protocol session counter; 'active-session-sctp': Active SCTP session counter; 'session-sctp': SCTP session counter; 'hitcount-timestamp': Last hit counts timestamp; 'rate-limit-drops': Rate Limit Drops;",
						},
					},
				},
			},
			"service_any": {
				Type: schema.TypeString, Optional: true, Default: "any", Description: "'any': any;",
			},
			"service_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocols": {
							Type: schema.TypeString, Optional: true, Description: "'tcp': tcp; 'udp': udp; 'sctp': sctp;",
						},
						"proto_id": {
							Type: schema.TypeInt, Optional: true, Description: "Protocol ID",
						},
						"obj_grp_service": {
							Type: schema.TypeString, Optional: true, Description: "service object group",
						},
						"icmp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ICMP",
						},
						"icmpv6": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ICMPv6",
						},
						"icmp_type": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP type number",
						},
						"special_type": {
							Type: schema.TypeString, Optional: true, Description: "'any-type': Any ICMP type; 'echo-reply': Type 0, echo reply; 'echo-request': Type 8, echo request; 'info-reply': Type 16, information reply; 'info-request': Type 15, information request; 'mask-reply': Type 18, address mask reply; 'mask-request': Type 17, address mask request; 'parameter-problem': Type 12, parameter problem; 'redirect': Type 5, redirect message; 'source-quench': Type 4, source quench; 'time-exceeded': Type 11, time exceeded; 'timestamp': Type 13, timestamp; 'timestamp-reply': Type 14, timestamp reply; 'dest-unreachable': Type 3, destination unreachable;",
						},
						"icmp_code": {
							Type: schema.TypeInt, Optional: true, Description: "ICMP code number",
						},
						"special_code": {
							Type: schema.TypeString, Optional: true, Description: "'any-code': Any ICMP code; 'frag-required': Code 4, fragmentation required; 'host-unreachable': Code 1, destination host unreachable; 'network-unreachable': Code 0, destination network unreachable; 'port-unreachable': Code 3, destination port unreachable; 'proto-unreachable': Code 2, destination protocol unreachable; 'route-failed': Code 5, source route failed;",
						},
						"icmpv6_type": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 type number",
						},
						"special_v6_type": {
							Type: schema.TypeString, Optional: true, Description: "'any-type': Any ICMPv6 type; 'dest-unreachable': Type 1, destination unreachable; 'echo-reply': Type 129, echo reply; 'echo-request': Type 128, echo request; 'packet-too-big': Type 2, packet too big; 'param-prob': Type 4, parameter problem; 'time-exceeded': Type 3, time exceeded;",
						},
						"icmpv6_code": {
							Type: schema.TypeInt, Optional: true, Description: "ICMPv6 code number",
						},
						"special_v6_code": {
							Type: schema.TypeString, Optional: true, Description: "'any-code': Any ICMPv6 code; 'addr-unreachable': Code 3, address unreachable; 'admin-prohibited': Code 1, admin prohibited; 'no-route': Code 0, no route to destination; 'not-neighbour': Code 2, not neighbor; 'port-unreachable': Code 4, destination port unreachable;",
						},
						"eq_src_port": {
							Type: schema.TypeInt, Optional: true, Description: "Equal to the port number",
						},
						"gt_src_port": {
							Type: schema.TypeInt, Optional: true, Description: "Greater than the port number",
						},
						"lt_src_port": {
							Type: schema.TypeInt, Optional: true, Description: "Lower than the port number",
						},
						"range_src_port": {
							Type: schema.TypeInt, Optional: true, Description: "Port range (Starting Port Number)",
						},
						"port_num_end_src": {
							Type: schema.TypeInt, Optional: true, Description: "Ending Port Number",
						},
						"eq_dst_port": {
							Type: schema.TypeInt, Optional: true, Description: "Equal to the port number",
						},
						"gt_dst_port": {
							Type: schema.TypeInt, Optional: true, Description: "Greater than the port number",
						},
						"lt_dst_port": {
							Type: schema.TypeInt, Optional: true, Description: "Lower than the port number",
						},
						"range_dst_port": {
							Type: schema.TypeInt, Optional: true, Description: "Port range (Starting Port Number)",
						},
						"port_num_end_dst": {
							Type: schema.TypeInt, Optional: true, Description: "Ending Port Number",
						},
						"sctp_template": {
							Type: schema.TypeString, Optional: true, Description: "SCTP Template",
						},
						"alg": {
							Type: schema.TypeString, Optional: true, Description: "'FTP': FTP; 'TFTP': TFTP; 'SIP': SIP; 'DNS': DNS; 'PPTP': PPTP; 'RTSP': RTSP; 'ESP': ESP;",
						},
					},
				},
			},
			"source_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_ip_subnet": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 IP Address",
						},
						"src_ipv6_subnet": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 IP Address",
						},
						"src_obj_network": {
							Type: schema.TypeString, Optional: true, Description: "Network object",
						},
						"src_obj_grp_network": {
							Type: schema.TypeString, Optional: true, Description: "Network object group",
						},
						"src_slb_server": {
							Type: schema.TypeString, Optional: true, Description: "SLB Real server name",
						},
					},
				},
			},
			"src_class_list": {
				Type: schema.TypeString, Optional: true, Description: "Match source IP against class-list",
			},
			"src_geoloc_list": {
				Type: schema.TypeString, Optional: true, Description: "Geolocation name list",
			},
			"src_geoloc_list_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use Geolocation list from shared partition",
			},
			"src_geoloc_name": {
				Type: schema.TypeString, Optional: true, Description: "Single geolocation name",
			},
			"src_ipv4_any": {
				Type: schema.TypeString, Optional: true, Default: "any", Description: "'any': Any IPv4 address;",
			},
			"src_ipv6_any": {
				Type: schema.TypeString, Optional: true, Default: "any", Description: "'any': Any IPv6 address;",
			},
			"src_threat_list": {
				Type: schema.TypeString, Optional: true, Description: "Bind threat-list for source IP based filtering",
			},
			"src_zone": {
				Type: schema.TypeString, Optional: true, Description: "Zone name",
			},
			"src_zone_any": {
				Type: schema.TypeString, Optional: true, Default: "any", Description: "'any': any;",
			},
			"status": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable rule; 'disable': Disable rule;",
			},
			"track_application": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable application statistic (functional only in action permit)",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vpn_ipsec_group_name": {
				Type: schema.TypeString, Optional: true, Description: "VPN IPsec Group name",
			},
			"vpn_ipsec_name": {
				Type: schema.TypeString, Optional: true, Description: "VPN IPsec name",
			},
		},
	}
}
func resourceRuleSetRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRuleSetRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceRuleSetRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRuleSetRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceRuleSetRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRuleSetRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectRuleSetRuleActionGroup1315(d []interface{}) edpt.RuleSetRuleActionGroup1315 {

	count1 := len(d)
	var ret edpt.RuleSetRuleActionGroup1315
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Type = in["type"].(string)
		ret.PermitLog = in["permit_log"].(int)
		ret.ResetLog = in["reset_log"].(int)
		ret.DenyLog = in["deny_log"].(int)
		ret.LoggingTemplateList = getSliceRuleSetRuleActionGroupLoggingTemplateList1316(in["logging_template_list"].([]interface{}))
		ret.ResetLogTemplateType = in["reset_log_template_type"].(string)
		ret.ResetFwLog = in["reset_fw_log"].(string)
		ret.DenyLogTemplateType = in["deny_log_template_type"].(string)
		ret.DenyFwLog = in["deny_fw_log"].(string)
		ret.ListenOnPort = in["listen_on_port"].(int)
		ret.Forward = in["forward"].(int)
		ret.Ipsec = in["ipsec"].(int)
		ret.IpsecGroup = in["ipsec_group"].(int)
		ret.VpnIpsecName = in["vpn_ipsec_name"].(string)
		ret.VpnIpsecGroupName = in["vpn_ipsec_group_name"].(string)
		ret.Cgnv6 = in["cgnv6"].(int)
		ret.Cgnv6Policy = in["cgnv6_policy"].(string)
		ret.Cgnv6LsnLid = in["cgnv6_lsn_lid"].(int)
		ret.Cgnv6DsLite = in["cgnv6_ds_lite"].(string)
		ret.Cgnv6DsLiteLsnLid = in["cgnv6_ds_lite_lsn_lid"].(int)
		ret.InspectPayload = in["inspect_payload"].(int)
		ret.PermitLimitPolicy = in["permit_limit_policy"].(int)
		ret.PermitRespondToUserMac = in["permit_respond_to_user_mac"].(int)
		ret.ResetRespondToUserMac = in["reset_respond_to_user_mac"].(int)
		ret.SetDscp = in["set_dscp"].(int)
		ret.DscpValue = in["dscp_value"].(string)
		ret.DscpNumber = in["dscp_number"].(int)
		//omit uuid
	}
	return ret
}

func getSliceRuleSetRuleActionGroupLoggingTemplateList1316(d []interface{}) []edpt.RuleSetRuleActionGroupLoggingTemplateList1316 {

	count1 := len(d)
	ret := make([]edpt.RuleSetRuleActionGroupLoggingTemplateList1316, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRuleActionGroupLoggingTemplateList1316
		oi.PermitLogTemplateType = in["permit_log_template_type"].(string)
		oi.PermitFwLog = in["permit_fw_log"].(string)
		oi.PermitCgnv6Log = in["permit_cgnv6_log"].(string)
		oi.PermitNetflowLog = in["permit_netflow_log"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetRuleAppList(d []interface{}) []edpt.RuleSetRuleAppList {

	count1 := len(d)
	ret := make([]edpt.RuleSetRuleAppList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRuleAppList
		oi.ObjGrpApplication = in["obj_grp_application"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.ProtocolTag = in["protocol_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetRuleDestList(d []interface{}) []edpt.RuleSetRuleDestList {

	count1 := len(d)
	ret := make([]edpt.RuleSetRuleDestList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRuleDestList
		oi.DstIpSubnet = in["dst_ip_subnet"].(string)
		oi.DstIpv6Subnet = in["dst_ipv6_subnet"].(string)
		oi.DstObjNetwork = in["dst_obj_network"].(string)
		oi.DstObjGrpNetwork = in["dst_obj_grp_network"].(string)
		oi.DstSlbServer = in["dst_slb_server"].(string)
		oi.DstSlbVserver = in["dst_slb_vserver"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetRuleDscpList(d []interface{}) []edpt.RuleSetRuleDscpList {

	count1 := len(d)
	ret := make([]edpt.RuleSetRuleDscpList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRuleDscpList
		oi.DscpValue = in["dscp_value"].(string)
		oi.DscpRangeStart = in["dscp_range_start"].(int)
		oi.DscpRangeEnd = in["dscp_range_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectRuleSetRuleMoveRule1317(d []interface{}) edpt.RuleSetRuleMoveRule1317 {

	count1 := len(d)
	var ret edpt.RuleSetRuleMoveRule1317
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Location = in["location"].(string)
		ret.TargetRule = in["target_rule"].(string)
	}
	return ret
}

func getSliceRuleSetRuleSamplingEnable(d []interface{}) []edpt.RuleSetRuleSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.RuleSetRuleSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRuleSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetRuleServiceList(d []interface{}) []edpt.RuleSetRuleServiceList {

	count1 := len(d)
	ret := make([]edpt.RuleSetRuleServiceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRuleServiceList
		oi.Protocols = in["protocols"].(string)
		oi.ProtoId = in["proto_id"].(int)
		oi.ObjGrpService = in["obj_grp_service"].(string)
		oi.Icmp = in["icmp"].(int)
		oi.Icmpv6 = in["icmpv6"].(int)
		oi.IcmpType = in["icmp_type"].(int)
		oi.SpecialType = in["special_type"].(string)
		oi.IcmpCode = in["icmp_code"].(int)
		oi.SpecialCode = in["special_code"].(string)
		oi.Icmpv6Type = in["icmpv6_type"].(int)
		oi.SpecialV6Type = in["special_v6_type"].(string)
		oi.Icmpv6Code = in["icmpv6_code"].(int)
		oi.SpecialV6Code = in["special_v6_code"].(string)
		oi.EqSrcPort = in["eq_src_port"].(int)
		oi.GtSrcPort = in["gt_src_port"].(int)
		oi.LtSrcPort = in["lt_src_port"].(int)
		oi.RangeSrcPort = in["range_src_port"].(int)
		oi.PortNumEndSrc = in["port_num_end_src"].(int)
		oi.EqDstPort = in["eq_dst_port"].(int)
		oi.GtDstPort = in["gt_dst_port"].(int)
		oi.LtDstPort = in["lt_dst_port"].(int)
		oi.RangeDstPort = in["range_dst_port"].(int)
		oi.PortNumEndDst = in["port_num_end_dst"].(int)
		oi.SctpTemplate = in["sctp_template"].(string)
		oi.Alg = in["alg"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRuleSetRuleSourceList(d []interface{}) []edpt.RuleSetRuleSourceList {

	count1 := len(d)
	ret := make([]edpt.RuleSetRuleSourceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRuleSourceList
		oi.SrcIpSubnet = in["src_ip_subnet"].(string)
		oi.SrcIpv6Subnet = in["src_ipv6_subnet"].(string)
		oi.SrcObjNetwork = in["src_obj_network"].(string)
		oi.SrcObjGrpNetwork = in["src_obj_grp_network"].(string)
		oi.SrcSlbServer = in["src_slb_server"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRuleSetRule(d *schema.ResourceData) edpt.RuleSetRule {
	var ret edpt.RuleSetRule
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ActionGroup = getObjectRuleSetRuleActionGroup1315(d.Get("action_group").([]interface{}))
	ret.Inst.AppList = getSliceRuleSetRuleAppList(d.Get("app_list").([]interface{}))
	ret.Inst.ApplicationAny = d.Get("application_any").(string)
	ret.Inst.Cgnv6DsLite = d.Get("cgnv6_ds_lite").(string)
	ret.Inst.Cgnv6DsLiteLog = d.Get("cgnv6_ds_lite_log").(int)
	ret.Inst.Cgnv6DsLiteLsnLid = d.Get("cgnv6_ds_lite_lsn_lid").(int)
	ret.Inst.Cgnv6FixedNatLog = d.Get("cgnv6_fixed_nat_log").(int)
	ret.Inst.Cgnv6Log = d.Get("cgnv6_log").(int)
	ret.Inst.Cgnv6LsnLid = d.Get("cgnv6_lsn_lid").(int)
	ret.Inst.Cgnv6LsnLog = d.Get("cgnv6_lsn_log").(int)
	ret.Inst.Cgnv6Policy = d.Get("cgnv6_policy").(string)
	ret.Inst.DestList = getSliceRuleSetRuleDestList(d.Get("dest_list").([]interface{}))
	ret.Inst.DscpList = getSliceRuleSetRuleDscpList(d.Get("dscp_list").([]interface{}))
	ret.Inst.DstClassList = d.Get("dst_class_list").(string)
	ret.Inst.DstDomainList = d.Get("dst_domain_list").(string)
	ret.Inst.DstGeolocList = d.Get("dst_geoloc_list").(string)
	ret.Inst.DstGeolocListShared = d.Get("dst_geoloc_list_shared").(int)
	ret.Inst.DstGeolocName = d.Get("dst_geoloc_name").(string)
	ret.Inst.DstIpv4Any = d.Get("dst_ipv4_any").(string)
	ret.Inst.DstIpv6Any = d.Get("dst_ipv6_any").(string)
	ret.Inst.DstThreatList = d.Get("dst_threat_list").(string)
	ret.Inst.DstZone = d.Get("dst_zone").(string)
	ret.Inst.DstZoneAny = d.Get("dst_zone_any").(string)
	ret.Inst.ForwardListenOnPort = d.Get("forward_listen_on_port").(int)
	ret.Inst.ForwardLog = d.Get("forward_log").(int)
	ret.Inst.FwLog = d.Get("fw_log").(int)
	ret.Inst.Fwlog = d.Get("fwlog").(int)
	ret.Inst.GtpTemplate = d.Get("gtp_template").(string)
	ret.Inst.IdleTimeout = d.Get("idle_timeout").(int)
	ret.Inst.InspectPayload = d.Get("inspect_payload").(int)
	ret.Inst.IpVersion = d.Get("ip_version").(string)
	ret.Inst.Lid = d.Get("lid").(int)
	ret.Inst.Lidlog = d.Get("lidlog").(int)
	ret.Inst.ListenOnPort = d.Get("listen_on_port").(int)
	ret.Inst.ListenOnPortLid = d.Get("listen_on_port_lid").(int)
	ret.Inst.ListenOnPortLidlog = d.Get("listen_on_port_lidlog").(int)
	ret.Inst.Log = d.Get("log").(int)
	ret.Inst.MoveRule = getObjectRuleSetRuleMoveRule1317(d.Get("move_rule").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Policy = d.Get("policy").(string)
	ret.Inst.Remark = d.Get("remark").(string)
	ret.Inst.ResetLid = d.Get("reset_lid").(int)
	ret.Inst.ResetLidlog = d.Get("reset_lidlog").(int)
	ret.Inst.SamplingEnable = getSliceRuleSetRuleSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServiceAny = d.Get("service_any").(string)
	ret.Inst.ServiceList = getSliceRuleSetRuleServiceList(d.Get("service_list").([]interface{}))
	ret.Inst.SourceList = getSliceRuleSetRuleSourceList(d.Get("source_list").([]interface{}))
	ret.Inst.SrcClassList = d.Get("src_class_list").(string)
	ret.Inst.SrcGeolocList = d.Get("src_geoloc_list").(string)
	ret.Inst.SrcGeolocListShared = d.Get("src_geoloc_list_shared").(int)
	ret.Inst.SrcGeolocName = d.Get("src_geoloc_name").(string)
	ret.Inst.SrcIpv4Any = d.Get("src_ipv4_any").(string)
	ret.Inst.SrcIpv6Any = d.Get("src_ipv6_any").(string)
	ret.Inst.SrcThreatList = d.Get("src_threat_list").(string)
	ret.Inst.SrcZone = d.Get("src_zone").(string)
	ret.Inst.SrcZoneAny = d.Get("src_zone_any").(string)
	ret.Inst.Status = d.Get("status").(string)
	ret.Inst.TrackApplication = d.Get("track_application").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.VpnIpsecGroupName = d.Get("vpn_ipsec_group_name").(string)
	ret.Inst.VpnIpsecName = d.Get("vpn_ipsec_name").(string)
	return ret
}
