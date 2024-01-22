package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_global`: Configure firewall parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceFwGlobalCreate,
		UpdateContext: resourceFwGlobalUpdate,
		ReadContext:   resourceFwGlobalRead,
		DeleteContext: resourceFwGlobalDelete,

		Schema: map[string]*schema.Schema{
			"alg_processing": {
				Type: schema.TypeString, Optional: true, Default: "honor-rule-set", Description: "'honor-rule-set': Honors firewall rule-sets (Default); 'override-rule-set': Override firewall rule-sets;",
			},
			"allow_non_syn_session_create": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow TCP non-syn packets to trigger session creation",
			},
			"disable_app_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disable_application_protocol": {
							Type: schema.TypeString, Optional: true, Description: "Disable specific application protocol",
						},
						"disable_application_category": {
							Type: schema.TypeString, Optional: true, Description: "'aaa': Protocol/application used for AAA (Authentification, Authorization and Accounting) purposes.; 'adult-content': Adult content protocol/application.; 'advertising': Advertising networks and applications.; 'application-enforcing-tls': Application known to enforce HSTS and thus use of TLS.; 'analytics-and-statistics': User analytics and statistics protocol/application.; 'anonymizers-and-proxies': Traffic-anonymization protocol/application.; 'audio-chat': Protocol/application used for Audio Chat.; 'basic': Covers all protocols required for basic classification, including most networking protocols as well as standard protocols like HTTP.; 'blog': Blogging platform protocol/application.; 'cdn': Protocol/application used for Content-Delivery Networks.; 'certification-authority': Certification Authority for SSL/TLS certificate.; 'chat': Protocol/application used for Text Chat.; 'classified-ads': Protocol/application used for Classified Advertisements.; 'cloud-based-services': SaaS and/or PaaS cloud based services.; 'crowdfunding': Service for funding a project or venture by raising small amounts of money from a large number of people, typically via the Internet.; 'cryptocurrency': Services for mining cryptocurrencies, for example a Crypto Web Browser (an application that mines crypto currency in the background while its user browses the web).; 'database': Database-specific protocols.; 'disposable-email': Service offering Disposable Email Accounts (DEA). DEA is a technique to share temporary email address between many users.; 'ebook-reader': Services for e-book readers, i.e. connected devices that display electronic books (typically using e-ink displays to reduce glare and eye strain).; 'education': Protocols offering education services and online courses.; 'email': Native email protocol.; 'enterprise': Protocol/application used in an enterprise network.; 'file-management': Protocol/application designed specifically for file management and exchange. This can include bona fide network protocols (like SMB) as well as web/cloud services (like Dropbox).; 'file-transfer': Protocol that offers file transferring as a secondary feature. This typically includes IM, WebMail, and other protocols that allow file transfers in addition to their principal function.; 'forum': Online forum protocol/application.; 'gaming': Protocol/application used by games.; 'healthcare': Protocols offering medical services, i.e protocols used in medical environment.; 'instant-messaging-and-multimedia-conferencing': Protocol/application used for Instant Messaging or Multi-Conferencing.; 'internet-of-things': Internet Of Things protocol/application.; 'map-service': Digital Maps service (web site and their related API).; 'mobile': Mobile-specific protocol/application.; 'multimedia-streaming': Protocol/application used for multimedia streaming.; 'networking': Protocol used for (inter) networking purpose.; 'news-portal': Protocol/application used for News Portals.; 'payment-service': Application offering online services for accepting electronic payments by a variety of payment methods (credit card, bank-based payments such as direct debit, bank transfer, etc).; 'peer-to-peer': Protocol/application used for Peer-to-peer purposes.; 'remote-access': Protocol/application used for remote access.; 'scada': SCADA (Supervisory control and data acquisition) protocols, all generations.; 'social-networks': Social networking application.; 'software-update': Auto-update protocol.; 'speedtest': Speedtest application allowing to access quality of Internet connection (upload, download, latency, etc).; 'standards-based': Protocol issued from standardized bodies such as IETF, ITU, IEEE, ETSI, OIF.; 'transportation': Transportation services, for example smartphone applications that allow users to hail a taxi.; 'video-chat': Protocol/application used for Video Chat.; 'voip': Application used for Voice-Over-IP.; 'vpn-tunnels': Protocol/application used for VPN or tunneling purposes.; 'web': Application based on HTTP/HTTPS.; 'web-e-commerce': Protocol/application used for E-commerce websites.; 'web-search-engines': Protocol/application used for Web search portals.; 'web-websites': Protocol/application used for Company Websites.; 'webmails': Web-based e-mail application.; 'web-ext-adult': Web Extension Adult; 'web-ext-auctions': Web Extension Auctions; 'web-ext-blogs': Web Extension Blogs; 'web-ext-business-and-economy': Web Extension Business and Economy; 'web-ext-cdns': Web Extension CDNs; 'web-ext-collaboration': Web Extension Collaboration; 'web-ext-computer-and-internet-info': Web Extension Computer and Internet Info; 'web-ext-computer-and-internet-security': Web Extension Computer and Internet Security; 'web-ext-dating': Web Extension Dating; 'web-ext-educational-institutions': Web Extension Educational Institutions; 'web-ext-entertainment-and-arts': Web Extension Entertainment and Arts; 'web-ext-fashion-and-beauty': Web Extension Fashion and Beauty; 'web-ext-file-share': Web Extension File Share; 'web-ext-financial-services': Web Extension Financial Services; 'web-ext-gambling': Web Extension Gambling; 'web-ext-games': Web Extension Games; 'web-ext-government': Web Extension Government; 'web-ext-health-and-medicine': Web Extension Health and Medicine; 'web-ext-individual-stock-advice-and-tools': Web Extension Individual Stock Advice and Tools; 'web-ext-internet-portals': Web Extension Internet Portals; 'web-ext-job-search': Web Extension Job Search; 'web-ext-local-information': Web Extension Local Information; 'web-ext-malware': Web Extension Malware; 'web-ext-motor-vehicles': Web Extension Motor Vehicles; 'web-ext-music': Web Extension Music; 'web-ext-news': Web Extension News; 'web-ext-p2p': Web Extension P2P; 'web-ext-parked-sites': Web Extension Parked Sites; 'web-ext-proxy-avoid-and-anonymizers': Web Extension Proxy Avoid and Anonymizers; 'web-ext-real-estate': Web Extension Real Estate; 'web-ext-reference-and-research': Web Extension Reference and Research; 'web-ext-search-engines': Web Extension Search Engines; 'web-ext-shopping': Web Extension Shopping; 'web-ext-social-network': Web Extension Social Network; 'web-ext-society': Web Extension Society; 'web-ext-software': Web Extension Software; 'web-ext-sports': Web Extension Sports; 'web-ext-streaming-media': Web Extension Streaming Media; 'web-ext-training-and-tools': Web Extension Training and Tools; 'web-ext-translation': Web Extension Translation; 'web-ext-travel': Web Extension Travel; 'web-ext-web-advertisements': Web Extension Web Advertisements; 'web-ext-web-based-email': Web Extension Web based Email; 'web-ext-web-hosting': Web Extension Web Hosting; 'web-ext-web-service': Web Extension Web Service;",
						},
					},
				},
			},
			"disable_application_metrics": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable exporting application protocol/category statistics to Harmony Controller",
			},
			"disable_ip_fw_sessions": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "disable create sessions for non TCP/UDP/ICMP",
			},
			"disable_undetermined_rule_logs": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "disable logs with undetermined rules",
			},
			"dsr_mode_support": {
				Type: schema.TypeString, Optional: true, Description: "'ipv4': support dsr for ipv4 traffic; 'ipv6': support dsr for ipv6 traffic; 'all': support dsr for both ipv4 and ipv6;",
			},
			"extended_matching": {
				Type: schema.TypeString, Optional: true, Description: "'disable': Disable extended matching;",
			},
			"inbound_refresh": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
			},
			"inbound_refresh_full_cone": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
			},
			"listen_on_port_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "STUN timeout (default: 2 minutes)",
			},
			"natip_ddos_protection": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable; 'disable': Disable;",
			},
			"permit_default_action": {
				Type: schema.TypeString, Optional: true, Description: "'forward': Forward; 'next-service-mode': Service to be applied chosen based on configuration;",
			},
			"respond_to_user_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the user's source MAC for the next hop rather than the routing table (default: off)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'tcp_fullcone_created': TCP Full-cone Created; 'tcp_fullcone_freed': TCP Full-cone Freed; 'udp_fullcone_created': UDP Full-cone Created; 'udp_fullcone_freed': UDP Full-cone Freed; 'fullcone_creation_failure': Full-Cone Creation Failure; 'data_session_created': Data Session Created; 'data_session_created_local': Data Session Created Local; 'dyn_blist_sess_sp': Dynamic Blacklist Session (Slowpath); 'data_session_freed': Data Session Freed; 'data_session_freed_local': Data Session Freed Local; 'dyn_blist_sess_created': Dynamic Blacklist Session Created; 'dyn_blist_sess_freed': Dynamic Blacklist Freed; 'dyn_blist_pkt_drop': Dynamic Blacklist - Packet Drop; 'dyn_blist_pkt_rate_low': Dynamic Blacklist - Pkt Rate Low; 'dyn_blist_pkt_rate_high': Dynamic Blacklist - Pkt Rate High; 'dyn_blist_version_mismatch': Dynamic Blacklist - Version Mismatch; 'dyn_blist_no_active_policy': Dynamic Blacklist - No Active Policy; 'fullcone_in_del_q': Full-cone session found in delete queue; 'fullcone_retry_lookup': Full-cone session retry look-up; 'fullcone_not_found': Full-cone session not found; 'fullcone_overflow_eim': Full-cone Session EIM Overflow; 'fullcone_overflow_eif': Full-cone Session EIF Overflow; 'udp_fullcone_created_shadow': Total UDP Full-cone sessions created; 'tcp_fullcone_created_shadow': Total TCP Full-cone sessions created; 'udp_fullcone_freed_shadow': Total UDP Full-cone sessions freed; 'tcp_fullcone_freed_shadow': Total TCP Full-cone sessions freed; 'fullcone_created': Total Full-cone sessions created; 'fullcone_freed': Total Full-cone sessions freed; 'fullcone_ext_too_many': Fullcone Extension Too Many; 'fullcone_ext_mem_allocated': Fullcone Extension Memory Allocated; 'fullcone_ext_mem_alloc_failure': Fullcone Extension Memory Allocate Failure; 'fullcone_ext_mem_alloc_init_faulure': Fullcone Extension Initialization Failure; 'fullcone_ext_mem_freed': Fullcone Extension Memory Freed; 'fullcone_ext_added': Fullcone Extension Added; 'ha_fullcone_failure': HA Full-cone Session Failure; 'data_session_created_shadow': Shadow Data Sessions Created; 'data_session_created_shadow_local': Shadow Data Sessions Created Local; 'data_session_freed_shadow': Shadow Data Sessions Freed; 'data_session_freed_shadow_local': Shadow Data Sessions Freed Local; 'active_fullcone_session': Total Active Full-cone sessions; 'limit-entry-failure': Limit Entry Creation Failure; 'limit-entry-allocated': Limit Entry Allocated; 'limit-entry-mem-freed': Limit Entry Freed; 'limit-entry-created': Limit Entry Created; 'limit-entry-found': Limit Entry Found; 'limit-entry-not-in-bucket': Limit Entry Not in Bucket; 'limit-entry-marked-deleted': Limit Entry Marked Deleted; 'undetermined-rule-counter': Undetermined rule detected; 'non_syn_pkt_fwd_allowed': Non-SYN pkt forward allowed; 'invalid-lid-drop': Invalid Lid Drop; 'src-session-limit-exceeded': Concurrent Session Limit Exceeded; 'uplink-pps-limit-exceeded': Uplink PPS Limit Exceeded; 'downlink-pps-limit-exceeded': Downlink PPS Limit Exceeded; 'total-pps-limit-exceeded': Total PPS Limit Exceeded; 'uplink-throughput-limit-exceeded': Uplink Throughput Limit Exceeded; 'downlink-throughput-limit-exceeded': Downlink Throughput Limit Exceeded; 'total-throughput-limit-exceeded': Total Throughput Limit Exceeded; 'cps-limit-exceeded': Connections Per Second Limit Exceeded; 'limit-exceeded': Per Second Limit Exceeded (Deprecated); 'limit-entry-per-cpu-mem-allocated': Limit Entry Memory Allocated (Deprecated); 'limit-entry-per-cpu-mem-allocation-failed': Limit Entry Memory Allocation Failed (Deprecated); 'limit-entry-per-cpu-mem-freed': Limit Entry Memory Freed (Deprecated); 'alg_default_port_disable': alg_default_port_disable; 'no_fwd_route': No Forward Route; 'no_rev_route': No Reverse Route; 'no_fwd_l2_dst': No Forward Mac Entry; 'no_rev_l2_dst': No Reverse Mac Entry; 'l2_dst_in_out_same': L2 route to same port as received; 'l2_vlan_changed': L2 forwarding vlan changed after session create; 'urpf_pkt_drop': URPF check packet drop; 'fwd_ingress_packets_tcp': Forward Ingress Packets TCP; 'fwd_egress_packets_tcp': Forward Egress Packets TCP; 'rev_ingress_packets_tcp': Reverse Ingress Packets TCP; 'rev_egress_packets_tcp': Reverse Egress Packets TCP; 'fwd_ingress_bytes_tcp': Forward Ingress Bytes TCP; 'fwd_egress_bytes_tcp': Forward Egress Bytes TCP; 'rev_ingress_bytes_tcp': Reverse Ingress Bytes TCP; 'rev_egress_bytes_tcp': Reverse Egress Bytes TCP; 'fwd_ingress_packets_udp': Forward Ingress Packets UDP; 'fwd_egress_packets_udp': Forward Egress Packets UDP; 'rev_ingress_packets_udp': Reverse Ingress Packets UDP; 'rev_egress_packets_udp': Reverse Egress Packets UDP; 'fwd_ingress_bytes_udp': Forward Ingress Bytes UDP; 'fwd_egress_bytes_udp': Forward Egress Bytes UDP; 'rev_ingress_bytes_udp': Reverse Ingress Bytes UDP; 'rev_egress_bytes_udp': Reverse Egress Bytes UDP; 'fwd_ingress_packets_icmp': Forward Ingress Packets ICMP; 'fwd_egress_packets_icmp': Forward Egress Packets ICMP; 'rev_ingress_packets_icmp': Reverse Ingress Packets ICMP; 'rev_egress_packets_icmp': Reverse Egress Packets ICMP; 'fwd_ingress_bytes_icmp': Forward Ingress Bytes ICMP; 'fwd_egress_bytes_icmp': Forward Egress Bytes ICMP; 'rev_ingress_bytes_icmp': Reverse Ingress Bytes ICMP; 'rev_egress_bytes_icmp': Reverse Egress Bytes ICMP; 'fwd_ingress_packets_others': Forward Ingress Packets OTHERS; 'fwd_egress_packets_others': Forward Egress Packets OTHERS; 'rev_ingress_packets_others': Reverse Ingress Packets OTHERS; 'rev_egress_packets_others': Reverse Egress Packets OTHERS; 'fwd_ingress_bytes_others': Forward Ingress Bytes OTHERS; 'fwd_egress_bytes_others': Forward Egress Bytes OTHERS; 'rev_ingress_bytes_others': Reverse Ingress Bytes OTHERS; 'rev_egress_bytes_others': Reverse Egress Bytes OTHERS; 'fwd_ingress_pkt_size_range1': Forward Ingress Packet size between 0 and 200; 'fwd_ingress_pkt_size_range2': Forward Ingress Packet size between 201 and 800; 'fwd_ingress_pkt_size_range3': Forward Ingress Packet size between 801 and 1550; 'fwd_ingress_pkt_size_range4': Forward Ingress Packet size between 1551 and 9000; 'fwd_egress_pkt_size_range1': Forward Egress Packet size between 0 and 200; 'fwd_egress_pkt_size_range2': Forward Egress Packet size between 201 and 800; 'fwd_egress_pkt_size_range3': Forward Egress Packet size between 801 and 1550; 'fwd_egress_pkt_size_range4': Forward Egress Packet size between 1551 and 9000; 'rev_ingress_pkt_size_range1': Reverse Ingress Packet size between 0 and 200; 'rev_ingress_pkt_size_range2': Reverse Ingress Packet size between 201 and 800; 'rev_ingress_pkt_size_range3': Reverse Ingress Packet size between 801 and 1550; 'rev_ingress_pkt_size_range4': Reverse Ingress Packet size between 1551 and 9000; 'rev_egress_pkt_size_range1': Reverse Egress Packet size between 0 and 200; 'rev_egress_pkt_size_range2': Reverse Egress Packet size between 201 and 800; 'rev_egress_pkt_size_range3': Reverse Egress Packet size between 801 and 1550; 'rev_egress_pkt_size_range4': Reverse Egress Packet size between 1551 and 9000;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceFwGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceFwGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceFwGlobalDisableAppList(d []interface{}) []edpt.FwGlobalDisableAppList {

	count1 := len(d)
	ret := make([]edpt.FwGlobalDisableAppList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwGlobalDisableAppList
		oi.DisableApplicationProtocol = in["disable_application_protocol"].(string)
		oi.DisableApplicationCategory = in["disable_application_category"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceFwGlobalSamplingEnable(d []interface{}) []edpt.FwGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.FwGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.FwGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointFwGlobal(d *schema.ResourceData) edpt.FwGlobal {
	var ret edpt.FwGlobal
	ret.Inst.AlgProcessing = d.Get("alg_processing").(string)
	ret.Inst.AllowNonSynSessionCreate = d.Get("allow_non_syn_session_create").(int)
	ret.Inst.DisableAppList = getSliceFwGlobalDisableAppList(d.Get("disable_app_list").([]interface{}))
	ret.Inst.DisableApplicationMetrics = d.Get("disable_application_metrics").(int)
	ret.Inst.DisableIpFwSessions = d.Get("disable_ip_fw_sessions").(int)
	ret.Inst.DisableUndeterminedRuleLogs = d.Get("disable_undetermined_rule_logs").(int)
	ret.Inst.DsrModeSupport = d.Get("dsr_mode_support").(string)
	ret.Inst.ExtendedMatching = d.Get("extended_matching").(string)
	ret.Inst.InboundRefresh = d.Get("inbound_refresh").(string)
	ret.Inst.InboundRefreshFullCone = d.Get("inbound_refresh_full_cone").(string)
	ret.Inst.ListenOnPortTimeout = d.Get("listen_on_port_timeout").(int)
	ret.Inst.NatipDdosProtection = d.Get("natip_ddos_protection").(string)
	ret.Inst.PermitDefaultAction = d.Get("permit_default_action").(string)
	ret.Inst.RespondToUserMac = d.Get("respond_to_user_mac").(int)
	ret.Inst.SamplingEnable = getSliceFwGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
