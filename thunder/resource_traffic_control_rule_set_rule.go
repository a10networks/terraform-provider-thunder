package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrafficControlRuleSetRule() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_traffic_control_rule_set_rule`: Configure traffic control rule\n\n__PLACEHOLDER__",
		CreateContext: resourceTrafficControlRuleSetRuleCreate,
		UpdateContext: resourceTrafficControlRuleSetRuleUpdate,
		ReadContext:   resourceTrafficControlRuleSetRuleRead,
		DeleteContext: resourceTrafficControlRuleSetRuleDelete,

		Schema: map[string]*schema.Schema{
			"action_group": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"limit_policy": {
							Type: schema.TypeInt, Optional: true, Description: "Limit policy Template",
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
			"ip_version": {
				Type: schema.TypeString, Optional: true, Default: "v4", Description: "'v4': IPv4 rule; 'v6': IPv6 rule;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Rule name",
			},
			"rule_set_name": {
				Type: schema.TypeString, Required: true, Description: "Rule set name",
			},
			"remark": {
				Type: schema.TypeString, Optional: true, Description: "Rule entry comment (Notes for this rule)",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'hit-count': Hit counts;",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceTrafficControlRuleSetRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRule(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTrafficControlRuleSetRuleRead(ctx, d, meta)
	}
	return diags
}

func resourceTrafficControlRuleSetRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRule(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTrafficControlRuleSetRuleRead(ctx, d, meta)
	}
	return diags
}
func resourceTrafficControlRuleSetRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRule(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTrafficControlRuleSetRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTrafficControlRuleSetRuleRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTrafficControlRuleSetRule(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectTrafficControlRuleSetRuleActionGroup1906(d []interface{}) edpt.TrafficControlRuleSetRuleActionGroup1906 {

	count1 := len(d)
	var ret edpt.TrafficControlRuleSetRuleActionGroup1906
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LimitPolicy = in["limit_policy"].(int)
		//omit uuid
	}
	return ret
}

func getSliceTrafficControlRuleSetRuleAppList(d []interface{}) []edpt.TrafficControlRuleSetRuleAppList {

	count1 := len(d)
	ret := make([]edpt.TrafficControlRuleSetRuleAppList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TrafficControlRuleSetRuleAppList
		oi.ObjGrpApplication = in["obj_grp_application"].(string)
		oi.Protocol = in["protocol"].(string)
		oi.ProtocolTag = in["protocol_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceTrafficControlRuleSetRuleDestList(d []interface{}) []edpt.TrafficControlRuleSetRuleDestList {

	count1 := len(d)
	ret := make([]edpt.TrafficControlRuleSetRuleDestList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TrafficControlRuleSetRuleDestList
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

func getSliceTrafficControlRuleSetRuleSamplingEnable(d []interface{}) []edpt.TrafficControlRuleSetRuleSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.TrafficControlRuleSetRuleSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TrafficControlRuleSetRuleSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceTrafficControlRuleSetRuleServiceList(d []interface{}) []edpt.TrafficControlRuleSetRuleServiceList {

	count1 := len(d)
	ret := make([]edpt.TrafficControlRuleSetRuleServiceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TrafficControlRuleSetRuleServiceList
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
		ret = append(ret, oi)
	}
	return ret
}

func getSliceTrafficControlRuleSetRuleSourceList(d []interface{}) []edpt.TrafficControlRuleSetRuleSourceList {

	count1 := len(d)
	ret := make([]edpt.TrafficControlRuleSetRuleSourceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TrafficControlRuleSetRuleSourceList
		oi.SrcIpSubnet = in["src_ip_subnet"].(string)
		oi.SrcIpv6Subnet = in["src_ipv6_subnet"].(string)
		oi.SrcObjNetwork = in["src_obj_network"].(string)
		oi.SrcObjGrpNetwork = in["src_obj_grp_network"].(string)
		oi.SrcSlbServer = in["src_slb_server"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointTrafficControlRuleSetRule(d *schema.ResourceData) edpt.TrafficControlRuleSetRule {
	var ret edpt.TrafficControlRuleSetRule
	ret.Inst.ActionGroup = getObjectTrafficControlRuleSetRuleActionGroup1906(d.Get("action_group").([]interface{}))
	ret.Inst.AppList = getSliceTrafficControlRuleSetRuleAppList(d.Get("app_list").([]interface{}))
	ret.Inst.ApplicationAny = d.Get("application_any").(string)
	ret.Inst.DestList = getSliceTrafficControlRuleSetRuleDestList(d.Get("dest_list").([]interface{}))
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
	ret.Inst.IpVersion = d.Get("ip_version").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RuleSetName = d.Get("rule_set_name").(string)
	ret.Inst.Remark = d.Get("remark").(string)
	ret.Inst.SamplingEnable = getSliceTrafficControlRuleSetRuleSamplingEnable(d.Get("sampling_enable").([]interface{}))
	ret.Inst.ServiceAny = d.Get("service_any").(string)
	ret.Inst.ServiceList = getSliceTrafficControlRuleSetRuleServiceList(d.Get("service_list").([]interface{}))
	ret.Inst.SourceList = getSliceTrafficControlRuleSetRuleSourceList(d.Get("source_list").([]interface{}))
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
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
