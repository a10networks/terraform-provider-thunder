package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosProtection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_protection`: DDOS protection\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosProtectionCreate,
		UpdateContext: resourceDdosProtectionUpdate,
		ReadContext:   resourceDdosProtectionRead,
		DeleteContext: resourceDdosProtectionDelete,

		Schema: map[string]*schema.Schema{
			"blacklist_reason_tracking": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable blacklist reason tracking",
			},
			"close_sess_for_unauth_src_without_rst": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "When closing unauthenticated sessions, don't send TCP RST for established TCP sessions. (Default disabled / sending TCP RST for",
			},
			"disable_advanced_core_analysis": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable advanced context info in coredump file",
			},
			"disable_delay_dynamic_src_learning": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable delay dynamic src entry learning",
			},
			"disable_on_reboot": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable DDoS protection upon reboot/reload",
			},
			"disallow_rst_ack_in_syn_auth": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disallow RST-ACK passing syn-auth",
			},
			"enable_now": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Override disable-on-reboot to enable runtime DDOS protection",
			},
			"fast_aging": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"half_open_conn_ratio": {
							Type: schema.TypeInt, Optional: true, Default: 25, Description: "Minimum half-open session to total session ratio before session fast aging will take effect (default 25)",
						},
						"half_open_conn_threshold": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Minimum half-open session (percentage) before session fast aging will take effect (default 1)",
						},
					},
				},
			},
			"fast_path_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable fast path in SLB processing",
			},
			"force_routing_on_transp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force use of routing in transparent mode",
			},
			"force_traffic_to_same_blade_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Allow traffic to be distributed among blades on Chassis",
			},
			"hw_blocking_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable hardware blacklist blocking for src or dst default entries (default disabled)",
			},
			"hw_blocking_threshold_limit": {
				Type: schema.TypeInt, Optional: true, Default: 10000, Description: "Threshold to initiate hardware blocking (default 10000)",
			},
			"ipv6_src_hash_mask_bits": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mask_bit_offset_1": {
							Type: schema.TypeInt, Optional: true, Description: "Configure mask bits",
						},
						"mask_bit_offset_2": {
							Type: schema.TypeInt, Optional: true, Description: "Configure mask bits",
						},
						"mask_bit_offset_3": {
							Type: schema.TypeInt, Optional: true, Description: "Configure mask bits",
						},
						"mask_bit_offset_4": {
							Type: schema.TypeInt, Optional: true, Description: "Configure mask bits",
						},
						"mask_bit_offset_5": {
							Type: schema.TypeInt, Optional: true, Description: "Configure mask bits",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"mpls": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable MPLS packet inspection",
			},
			"multi_pu_zone_distribution": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"regular_rebalance": {
							Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': enable; 'disable': disable;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"non_zero_win_size_syncookie": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send syn-cookie with fix TCP window size if SYN packet has zero window size  (default disabled)",
			},
			"per_service_szp_entry_limit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_tcp_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range dns-tcp",
						},
						"dns_udp_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range dns-udp",
						},
						"http_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range http",
						},
						"tcp_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range tcp",
						},
						"udp_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range udp",
						},
						"ssl_l4_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range ssl-l4",
						},
						"sip_udp_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range sip-udp",
						},
						"sip_tcp_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range sip-tcp",
						},
						"quic_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for port / port-range quic",
						},
						"ip_proto_icmp_v4_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for ip-proto icmp-v4",
						},
						"ip_proto_icmp_v6_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for ip-proto icmp-v6",
						},
						"ip_proto_other_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for ip-proto other",
						},
						"ip_proto_gre_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for ip-proto gre",
						},
						"ip_proto_ipv4_encap_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for ip-proto ipv4-encap",
						},
						"ip_proto_ipv6_encap_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for ip-proto ipv6-encap",
						},
						"ip_proto_custom_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Szp limit for custom ip-proto",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"pkt_rate_limit_on_reassemble": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': enable; 'disable': disable (Default);",
			},
			"progression_tracking": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable; 'disable': disable;",
			},
			"rate_interval": {
				Type: schema.TypeString, Optional: true, Default: "100ms", Description: "'100ms': 100ms; '1sec': 1sec;",
			},
			"rate_limit_sync_interval": {
				Type: schema.TypeInt, Optional: true, Default: 3, Description: "Multi-PU rate limit syncing interval (default 3)",
			},
			"rexmit_syn_log": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ddos per flow rexmit syn exceeded log",
			},
			"src_dst_entry_limit": {
				Type: schema.TypeString, Optional: true, Default: "16M", Description: "'8M': 8 Million; '16M': 16 Million; 'unlimited': Unlimited; 'platform-default': Half of platform maximum;",
			},
			"src_hash_function": {
				Type: schema.TypeString, Optional: true, Default: "v1", Description: "'v1': v1; 'v2': v2;",
			},
			"src_ip_hash_bit": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Configure which bit hashed on",
			},
			"src_ipv6_hash_bit": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Configure which bit hashed on",
			},
			"src_zone_port_entry_limit": {
				Type: schema.TypeString, Optional: true, Default: "16M", Description: "'8M': 8 Million; '16M': 16 Million; 'unlimited': Unlimited; 'platform-default': Half of platform maximum;",
			},
			"szp_clist_warn_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Set threshold percentage of \"max-src-dst-entry\" for generating warning logs. Including start and end.",
			},
			"szp_warn_exceed_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send logs if src-zone-port count exceeds \"max-src-dst-entry\"",
			},
			"szp_warn_threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Set threshold percentage of \"max-src-dst-entry\" for generating warning logs. Including start and end.",
			},
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': enable; 'disable': disable;",
			},
			"use_route": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use route table, default use receive hop for device initiated traffic",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vxlan_outbound_check": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': enable; 'disable': disable;",
			},
		},
	}
}
func resourceDdosProtectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosProtectionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosProtectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosProtectionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosProtectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosProtectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosProtectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosProtection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosProtectionFastAging(d []interface{}) edpt.DdosProtectionFastAging {

	count1 := len(d)
	var ret edpt.DdosProtectionFastAging
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HalfOpenConnRatio = in["half_open_conn_ratio"].(int)
		ret.HalfOpenConnThreshold = in["half_open_conn_threshold"].(int)
	}
	return ret
}

func getObjectDdosProtectionIpv6SrcHashMaskBits330(d []interface{}) edpt.DdosProtectionIpv6SrcHashMaskBits330 {

	count1 := len(d)
	var ret edpt.DdosProtectionIpv6SrcHashMaskBits330
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaskBitOffset1 = in["mask_bit_offset_1"].(int)
		ret.MaskBitOffset2 = in["mask_bit_offset_2"].(int)
		ret.MaskBitOffset3 = in["mask_bit_offset_3"].(int)
		ret.MaskBitOffset4 = in["mask_bit_offset_4"].(int)
		ret.MaskBitOffset5 = in["mask_bit_offset_5"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDdosProtectionMultiPuZoneDistribution331(d []interface{}) edpt.DdosProtectionMultiPuZoneDistribution331 {

	count1 := len(d)
	var ret edpt.DdosProtectionMultiPuZoneDistribution331
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RegularRebalance = in["regular_rebalance"].(string)
		//omit uuid
	}
	return ret
}

func getObjectDdosProtectionPerServiceSzpEntryLimit332(d []interface{}) edpt.DdosProtectionPerServiceSzpEntryLimit332 {

	count1 := len(d)
	var ret edpt.DdosProtectionPerServiceSzpEntryLimit332
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DnsTcpLimit = in["dns_tcp_limit"].(int)
		ret.DnsUdpLimit = in["dns_udp_limit"].(int)
		ret.HttpLimit = in["http_limit"].(int)
		ret.TcpLimit = in["tcp_limit"].(int)
		ret.UdpLimit = in["udp_limit"].(int)
		ret.SslL4Limit = in["ssl_l4_limit"].(int)
		ret.SipUdpLimit = in["sip_udp_limit"].(int)
		ret.SipTcpLimit = in["sip_tcp_limit"].(int)
		ret.QuicLimit = in["quic_limit"].(int)
		ret.IpProtoIcmpV4Limit = in["ip_proto_icmp_v4_limit"].(int)
		ret.IpProtoIcmpV6Limit = in["ip_proto_icmp_v6_limit"].(int)
		ret.IpProtoOtherLimit = in["ip_proto_other_limit"].(int)
		ret.IpProtoGreLimit = in["ip_proto_gre_limit"].(int)
		ret.IpProtoIpv4EncapLimit = in["ip_proto_ipv4_encap_limit"].(int)
		ret.IpProtoIpv6EncapLimit = in["ip_proto_ipv6_encap_limit"].(int)
		ret.IpProtoCustomLimit = in["ip_proto_custom_limit"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointDdosProtection(d *schema.ResourceData) edpt.DdosProtection {
	var ret edpt.DdosProtection
	ret.Inst.BlacklistReasonTracking = d.Get("blacklist_reason_tracking").(int)
	ret.Inst.CloseSessForUnauthSrcWithoutRst = d.Get("close_sess_for_unauth_src_without_rst").(int)
	ret.Inst.DisableAdvancedCoreAnalysis = d.Get("disable_advanced_core_analysis").(int)
	ret.Inst.DisableDelayDynamicSrcLearning = d.Get("disable_delay_dynamic_src_learning").(int)
	ret.Inst.DisableOnReboot = d.Get("disable_on_reboot").(int)
	ret.Inst.DisallowRstAckInSynAuth = d.Get("disallow_rst_ack_in_syn_auth").(int)
	ret.Inst.EnableNow = d.Get("enable_now").(int)
	ret.Inst.FastAging = getObjectDdosProtectionFastAging(d.Get("fast_aging").([]interface{}))
	ret.Inst.FastPathDisable = d.Get("fast_path_disable").(int)
	ret.Inst.ForceRoutingOnTransp = d.Get("force_routing_on_transp").(int)
	ret.Inst.ForceTrafficToSameBladeDisable = d.Get("force_traffic_to_same_blade_disable").(int)
	ret.Inst.HwBlockingEnable = d.Get("hw_blocking_enable").(int)
	ret.Inst.HwBlockingThresholdLimit = d.Get("hw_blocking_threshold_limit").(int)
	ret.Inst.Ipv6SrcHashMaskBits = getObjectDdosProtectionIpv6SrcHashMaskBits330(d.Get("ipv6_src_hash_mask_bits").([]interface{}))
	ret.Inst.Mpls = d.Get("mpls").(int)
	ret.Inst.MultiPuZoneDistribution = getObjectDdosProtectionMultiPuZoneDistribution331(d.Get("multi_pu_zone_distribution").([]interface{}))
	ret.Inst.NonZeroWinSizeSyncookie = d.Get("non_zero_win_size_syncookie").(int)
	ret.Inst.PerServiceSzpEntryLimit = getObjectDdosProtectionPerServiceSzpEntryLimit332(d.Get("per_service_szp_entry_limit").([]interface{}))
	ret.Inst.PktRateLimitOnReassemble = d.Get("pkt_rate_limit_on_reassemble").(string)
	ret.Inst.ProgressionTracking = d.Get("progression_tracking").(string)
	ret.Inst.RateInterval = d.Get("rate_interval").(string)
	ret.Inst.RateLimitSyncInterval = d.Get("rate_limit_sync_interval").(int)
	ret.Inst.RexmitSynLog = d.Get("rexmit_syn_log").(int)
	ret.Inst.SrcDstEntryLimit = d.Get("src_dst_entry_limit").(string)
	ret.Inst.SrcHashFunction = d.Get("src_hash_function").(string)
	ret.Inst.SrcIpHashBit = d.Get("src_ip_hash_bit").(int)
	ret.Inst.SrcIpv6HashBit = d.Get("src_ipv6_hash_bit").(int)
	ret.Inst.SrcZonePortEntryLimit = d.Get("src_zone_port_entry_limit").(string)
	ret.Inst.SzpClistWarnThreshold = d.Get("szp_clist_warn_threshold").(int)
	ret.Inst.SzpWarnExceedEnable = d.Get("szp_warn_exceed_enable").(int)
	ret.Inst.SzpWarnThreshold = d.Get("szp_warn_threshold").(int)
	ret.Inst.Toggle = d.Get("toggle").(string)
	ret.Inst.UseRoute = d.Get("use_route").(int)
	//omit uuid
	ret.Inst.VxlanOutboundCheck = d.Get("vxlan_outbound_check").(string)
	return ret
}
