package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSessionFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_session_filter`: Create a convenience Filter to display/clear sessions\n\n__PLACEHOLDER__",
		CreateContext: resourceSessionFilterCreate,
		UpdateContext: resourceSessionFilterUpdate,
		ReadContext:   resourceSessionFilterRead,
		DeleteContext: resourceSessionFilterDelete,

		Schema: map[string]*schema.Schema{
			"filter_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"session_type": {
							Type: schema.TypeString, Optional: true, Description: "'ipv6': Display ipv6 sessions only; 'sip': SIP sessions;",
						},
						"source_addr": {
							Type: schema.TypeString, Optional: true, Description: "Forward Source IP (Source IP address)",
						},
						"source_mask": {
							Type: schema.TypeString, Optional: true, Description: "Forward Source IP Subnet (Source Netmask)",
						},
						"source_port": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Source Port",
						},
						"dest_addr": {
							Type: schema.TypeString, Optional: true, Description: "Forward Destination IP (Destination IP address)",
						},
						"dest_mask": {
							Type: schema.TypeString, Optional: true, Description: "Forward Destination IP Subnet (Destination Netmask)",
						},
						"dport2": {
							Type: schema.TypeInt, Optional: true, Description: "Forward Destination Port (Dest Port)",
						},
						"app": {
							Type: schema.TypeString, Optional: true, Description: "Specify application(s), separated by comma (For example: http,tcp)",
						},
						"app_category": {
							Type: schema.TypeString, Optional: true, Description: "'aaa': Protocol/application used for AAA (Authentification, Authorization and Accounting) purposes.; 'adult-content': Adult content protocol/application.; 'advertising': Advertising networks and applications.; 'application-enforcing-tls': Application known to enforce HSTS and thus use of TLS.; 'analytics-and-statistics': User analytics and statistics protocol/application.; 'anonymizers-and-proxies': Traffic-anonymization protocol/application.; 'audio-chat': Protocol/application used for Audio Chat.; 'basic': Covers all protocols required for basic classification, including most networking protocols as well as standard protocols like HTTP.; 'blog': Blogging platform protocol/application.; 'cdn': Protocol/application used for Content-Delivery Networks.; 'certification-authority': Certification Authority for SSL/TLS certificate.; 'chat': Protocol/application used for Text Chat.; 'classified-ads': Protocol/application used for Classified Advertisements.; 'cloud-based-services': SaaS and/or PaaS cloud based services.; 'crowdfunding': Service for funding a project or venture by raising small amounts of money from a large number of people, typically via the Internet.; 'cryptocurrency': Services for mining cryptocurrencies, for example a Crypto Web Browser (an application that mines crypto currency in the background while its user browses the web).; 'database': Database-specific protocols.; 'disposable-email': Service offering Disposable Email Accounts (DEA). DEA is a technique to share temporary email address between many users.; 'ebook-reader': Services for e-book readers, i.e. connected devices that display electronic books (typically using e-ink displays to reduce glare and eye strain).; 'education': Protocols offering education services and online courses.; 'email': Native email protocol.; 'enterprise': Protocol/application used in an enterprise network.; 'file-management': Protocol/application designed specifically for file management and exchange. This can include bona fide network protocols (like SMB) as well as web/cloud services (like Dropbox).; 'file-transfer': Protocol that offers file transferring as a secondary feature. This typically includes IM, WebMail, and other protocols that allow file transfers in addition to their principal function.; 'forum': Online forum protocol/application.; 'gaming': Protocol/application used by games.; 'healthcare': Protocols offering medical services, i.e protocols used in medical environment.; 'instant-messaging-and-multimedia-conferencing': Protocol/application used for Instant Messaging or Multi-Conferencing.; 'internet-of-things': Internet Of Things protocol/application.; 'map-service': Digital Maps service (web site and their related API).; 'mobile': Mobile-specific protocol/application.; 'multimedia-streaming': Protocol/application used for multimedia streaming.; 'networking': Protocol used for (inter) networking purpose.; 'news-portal': Protocol/application used for News Portals.; 'payment-service': Application offering online services for accepting electronic payments by a variety of payment methods (credit card, bank-based payments such as direct debit, bank transfer, etc).; 'peer-to-peer': Protocol/application used for Peer-to-peer purposes.; 'remote-access': Protocol/application used for remote access.; 'scada': SCADA (Supervisory control and data acquisition) protocols, all generations.; 'social-networks': Social networking application.; 'software-update': Auto-update protocol.; 'speedtest': Speedtest application allowing to access quality of Internet connection (upload, download, latency, etc).; 'standards-based': Protocol issued from standardized bodies such as IETF, ITU, IEEE, ETSI, OIF.; 'transportation': Transportation services, for example smartphone applications that allow users to hail a taxi.; 'video-chat': Protocol/application used for Video Chat.; 'voip': Application used for Voice-Over-IP.; 'vpn-tunnels': Protocol/application used for VPN or tunneling purposes.; 'web': Application based on HTTP/HTTPS.; 'web-e-commerce': Protocol/application used for E-commerce websites.; 'web-search-engines': Protocol/application used for Web search portals.; 'web-websites': Protocol/application used for Company Websites.; 'webmails': Web-based e-mail application.; 'web-ext-adult': Web Extension Adult; 'web-ext-auctions': Web Extension Auctions; 'web-ext-blogs': Web Extension Blogs; 'web-ext-business-and-economy': Web Extension Business and Economy; 'web-ext-cdns': Web Extension CDNs; 'web-ext-collaboration': Web Extension Collaboration; 'web-ext-computer-and-internet-info': Web Extension Computer and Internet Info; 'web-ext-computer-and-internet-security': Web Extension Computer and Internet Security; 'web-ext-dating': Web Extension Dating; 'web-ext-educational-institutions': Web Extension Educational Institutions; 'web-ext-entertainment-and-arts': Web Extension Entertainment and Arts; 'web-ext-fashion-and-beauty': Web Extension Fashion and Beauty; 'web-ext-file-share': Web Extension File Share; 'web-ext-financial-services': Web Extension Financial Services; 'web-ext-gambling': Web Extension Gambling; 'web-ext-games': Web Extension Games; 'web-ext-government': Web Extension Government; 'web-ext-health-and-medicine': Web Extension Health and Medicine; 'web-ext-individual-stock-advice-and-tools': Web Extension Individual Stock Advice and Tools; 'web-ext-internet-portals': Web Extension Internet Portals; 'web-ext-job-search': Web Extension Job Search; 'web-ext-local-information': Web Extension Local Information; 'web-ext-malware': Web Extension Malware; 'web-ext-motor-vehicles': Web Extension Motor Vehicles; 'web-ext-music': Web Extension Music; 'web-ext-news': Web Extension News; 'web-ext-p2p': Web Extension P2P; 'web-ext-parked-sites': Web Extension Parked Sites; 'web-ext-proxy-avoid-and-anonymizers': Web Extension Proxy Avoid and Anonymizers; 'web-ext-real-estate': Web Extension Real Estate; 'web-ext-reference-and-research': Web Extension Reference and Research; 'web-ext-search-engines': Web Extension Search Engines; 'web-ext-shopping': Web Extension Shopping; 'web-ext-social-network': Web Extension Social Network; 'web-ext-society': Web Extension Society; 'web-ext-software': Web Extension Software; 'web-ext-sports': Web Extension Sports; 'web-ext-streaming-media': Web Extension Streaming Media; 'web-ext-training-and-tools': Web Extension Training and Tools; 'web-ext-translation': Web Extension Translation; 'web-ext-travel': Web Extension Travel; 'web-ext-web-advertisements': Web Extension Web Advertisements; 'web-ext-web-based-email': Web Extension Web based Email; 'web-ext-web-hosting': Web Extension Web Hosting; 'web-ext-web-service': Web Extension Web Service;",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Session filter name",
			},
			"set": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set filter criteria",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSessionFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessionFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSessionFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceSessionFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessionFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSessionFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceSessionFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessionFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSessionFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSessionFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSessionFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSessionFilterFilterCfg(d []interface{}) edpt.SessionFilterFilterCfg {

	count1 := len(d)
	var ret edpt.SessionFilterFilterCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SessionType = in["session_type"].(string)
		ret.SourceAddr = in["source_addr"].(string)
		ret.SourceMask = in["source_mask"].(string)
		ret.SourcePort = in["source_port"].(int)
		ret.DestAddr = in["dest_addr"].(string)
		ret.DestMask = in["dest_mask"].(string)
		ret.Dport2 = in["dport2"].(int)
		ret.App = in["app"].(string)
		ret.AppCategory = in["app_category"].(string)
	}
	return ret
}

func dataToEndpointSessionFilter(d *schema.ResourceData) edpt.SessionFilter {
	var ret edpt.SessionFilter
	ret.Inst.FilterCfg = getObjectSessionFilterFilterCfg(d.Get("filter_cfg").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Set = d.Get("set").(int)
	//omit uuid
	return ret
}
