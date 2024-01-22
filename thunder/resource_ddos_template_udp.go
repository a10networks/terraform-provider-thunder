package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosTemplateUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_template_udp`: UDP template configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosTemplateUdpCreate,
		UpdateContext: resourceDdosTemplateUdpUpdate,
		ReadContext:   resourceDdosTemplateUdpRead,
		DeleteContext: resourceDdosTemplateUdpDelete,

		Schema: map[string]*schema.Schema{
			"age": {
				Type: schema.TypeInt, Optional: true, Description: "Configure session age(in minutes) for UDP sessions",
			},
			"drop_known_resp_src_port_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"drop_known_resp_src_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop well-known if src-port is less than 1024",
						},
						"exclude_src_resp_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "excluding src port equal destination port",
						},
					},
				},
			},
			"drop_ntp_monlist": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Drop NTP monlist request/response",
			},
			"filter_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"udp_filter_seq": {
							Type: schema.TypeInt, Required: true, Description: "Sequence number",
						},
						"udp_filter_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"byte_offset_filter": {
							Type: schema.TypeString, Optional: true, Description: "Filter Expression using Berkeley Packet Filter syntax",
						},
						"udp_filter_unmatched": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "action taken when it does not match",
						},
						"udp_filter_action": {
							Type: schema.TypeString, Optional: true, Description: "'blacklist-src': Also blacklist the source when action is taken; 'whitelist-src': Whitelist the source after filter passes, packets are dropped until then; 'count-only': Take no action and continue processing the next filter;",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"max_payload_size": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum UDP payload size for each single packet",
			},
			"min_payload_size": {
				Type: schema.TypeInt, Optional: true, Description: "Minimum UDP payload size for each single packet",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DDOS UDP Template Name",
			},
			"per_conn_pkt_rate_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Packet rate limit per connection per rate-interval",
			},
			"per_conn_rate_interval": {
				Type: schema.TypeString, Optional: true, Default: "1sec", Description: "'100ms': 100ms; '1sec': 1sec;",
			},
			"previous_salt_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Token-Authentication previous salt-prefix timeout in minutes, default is 1 min",
			},
			"public_ipv4_addr": {
				Type: schema.TypeString, Optional: true, Description: "IP address",
			},
			"public_ipv6_addr": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 address",
			},
			"spoof_detect_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"spoof_detect": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Force client to retry on udp",
						},
						"min_retry_gap_interval": {
							Type: schema.TypeString, Optional: true, Default: "1sec", Description: "'100ms': 100ms; '1sec': 1sec;",
						},
						"spoof_detect_retry_timeout_val_only": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "timeout in seconds",
						},
						"min_retry_gap": {
							Type: schema.TypeInt, Optional: true, Description: "Optional minimum gap between 2 UDP packets for spoof-detect pass, unit is specified by min-retry-gap-interval",
						},
						"spoof_detect_retry_timeout": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "timeout in seconds",
						},
					},
				},
			},
			"token_authentication": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Token Authentication",
			},
			"token_authentication_formula": {
				Type: schema.TypeString, Optional: true, Description: "'md5_Salt-SrcIp-SrcPort-DstIp-DstPort': md5 of Salt-SrcIp-SrcPort-DstIp-DstPort; 'md5_Salt-DstIp-DstPort': md5 of Salt-DstIp-DstPort; 'md5_Salt-SrcIp-DstIp': md5 of Salt-SrcIp-DstIp; 'md5_Salt-SrcPort-DstPort': md5 of Salt-SrcPort-DstPort; 'md5_Salt-UintDstIp-DstPort': Using the uint value of IP for md5 of Salt-DstIp-DstPort; 'sha1_Salt-SrcIp-SrcPort-DstIp-DstPort': sha1 of Salt-SrcIp-SrcPort-DstIp-DstPort; 'sha1_Salt-DstIp-DstPort': sha1 of Salt-DstIp-DstPort; 'sha1_Salt-SrcIp-DstIp': sha1 of Salt-SrcIp-DstIp; 'sha1_Salt-SrcPort-DstPort': sha1 of Salt-SrcPort-DstPort; 'sha1_Salt-UintDstIp-DstPort': Using the uint value of IP for sha1 of Salt-DstIp-DstPort;",
			},
			"token_authentication_hw_assist_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "token-authentication disable hardware assistance",
			},
			"token_authentication_public_address": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "The server public IP address",
			},
			"token_authentication_salt_prefix": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "token-authentication salt-prefix",
			},
			"token_authentication_salt_prefix_curr": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"token_authentication_salt_prefix_prev": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"tunnel_encap": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_encap": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Tunnel encapsulation using IP in IP",
						},
						"always": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_addr": {
										Type: schema.TypeString, Optional: true, Description: "IPv4 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
									},
									"preserve_src_ipv4": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use original source ip for encapsulation",
									},
									"ipv6_addr": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
									},
									"preserve_src_ipv6": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use original source ip for encapsulation",
									},
								},
							},
						},
						"gre_encap": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Tunnel encapsulation using GRE",
						},
						"gre_always": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gre_ipv4": {
										Type: schema.TypeString, Optional: true, Description: "IPv4 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
									},
									"key_ipv4": {
										Type: schema.TypeString, Optional: true, Description: "Encapsulate with key (Hexadecimal 0x0-0xFFFFFFFF,decimal 0-4294967295)",
									},
									"preserve_src_ipv4_gre": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use original source ip for encapsulation",
									},
									"gre_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "IPv6 address (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
									},
									"key_ipv6": {
										Type: schema.TypeString, Optional: true, Description: "Encapsulate with key (Hexadecimal 0x0-0xFFFFFFFF,decimal 0-4294967295)",
									},
									"preserve_src_ipv6_gre": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use original source ip for encapsulation",
									},
								},
							},
						},
					},
				},
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
func resourceDdosTemplateUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosTemplateUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosTemplateUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosTemplateUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosTemplateUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosTemplateUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosTemplateUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosTemplateUdpDropKnownRespSrcPortCfg(d []interface{}) edpt.DdosTemplateUdpDropKnownRespSrcPortCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateUdpDropKnownRespSrcPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DropKnownRespSrcPort = in["drop_known_resp_src_port"].(int)
		ret.ExcludeSrcRespPort = in["exclude_src_resp_port"].(int)
	}
	return ret
}

func getSliceDdosTemplateUdpFilterList(d []interface{}) []edpt.DdosTemplateUdpFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosTemplateUdpFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosTemplateUdpFilterList
		oi.UdpFilterSeq = in["udp_filter_seq"].(int)
		oi.UdpFilterRegex = in["udp_filter_regex"].(string)
		oi.ByteOffsetFilter = in["byte_offset_filter"].(string)
		oi.UdpFilterUnmatched = in["udp_filter_unmatched"].(int)
		oi.UdpFilterAction = in["udp_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosTemplateUdpSpoofDetectCfg(d []interface{}) edpt.DdosTemplateUdpSpoofDetectCfg {

	count1 := len(d)
	var ret edpt.DdosTemplateUdpSpoofDetectCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SpoofDetect = in["spoof_detect"].(int)
		ret.MinRetryGapInterval = in["min_retry_gap_interval"].(string)
		ret.SpoofDetectRetryTimeoutValOnly = in["spoof_detect_retry_timeout_val_only"].(int)
		ret.MinRetryGap = in["min_retry_gap"].(int)
		ret.SpoofDetectRetryTimeout = in["spoof_detect_retry_timeout"].(int)
	}
	return ret
}

func getObjectDdosTemplateUdpTunnelEncap(d []interface{}) edpt.DdosTemplateUdpTunnelEncap {

	count1 := len(d)
	var ret edpt.DdosTemplateUdpTunnelEncap
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpEncap = in["ip_encap"].(int)
		ret.Always = getObjectDdosTemplateUdpTunnelEncapAlways(in["always"].([]interface{}))
		ret.GreEncap = in["gre_encap"].(int)
		ret.GreAlways = getObjectDdosTemplateUdpTunnelEncapGreAlways(in["gre_always"].([]interface{}))
	}
	return ret
}

func getObjectDdosTemplateUdpTunnelEncapAlways(d []interface{}) edpt.DdosTemplateUdpTunnelEncapAlways {

	count1 := len(d)
	var ret edpt.DdosTemplateUdpTunnelEncapAlways
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4Addr = in["ipv4_addr"].(string)
		ret.PreserveSrcIpv4 = in["preserve_src_ipv4"].(int)
		ret.Ipv6Addr = in["ipv6_addr"].(string)
		ret.PreserveSrcIpv6 = in["preserve_src_ipv6"].(int)
	}
	return ret
}

func getObjectDdosTemplateUdpTunnelEncapGreAlways(d []interface{}) edpt.DdosTemplateUdpTunnelEncapGreAlways {

	count1 := len(d)
	var ret edpt.DdosTemplateUdpTunnelEncapGreAlways
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GreIpv4 = in["gre_ipv4"].(string)
		ret.KeyIpv4 = in["key_ipv4"].(string)
		ret.PreserveSrcIpv4Gre = in["preserve_src_ipv4_gre"].(int)
		ret.GreIpv6 = in["gre_ipv6"].(string)
		ret.KeyIpv6 = in["key_ipv6"].(string)
		ret.PreserveSrcIpv6Gre = in["preserve_src_ipv6_gre"].(int)
	}
	return ret
}

func dataToEndpointDdosTemplateUdp(d *schema.ResourceData) edpt.DdosTemplateUdp {
	var ret edpt.DdosTemplateUdp
	ret.Inst.Age = d.Get("age").(int)
	ret.Inst.DropKnownRespSrcPortCfg = getObjectDdosTemplateUdpDropKnownRespSrcPortCfg(d.Get("drop_known_resp_src_port_cfg").([]interface{}))
	ret.Inst.DropNtpMonlist = d.Get("drop_ntp_monlist").(int)
	ret.Inst.FilterList = getSliceDdosTemplateUdpFilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.MaxPayloadSize = d.Get("max_payload_size").(int)
	ret.Inst.MinPayloadSize = d.Get("min_payload_size").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PerConnPktRateLimit = d.Get("per_conn_pkt_rate_limit").(int)
	ret.Inst.PerConnRateInterval = d.Get("per_conn_rate_interval").(string)
	ret.Inst.PreviousSaltTimeout = d.Get("previous_salt_timeout").(int)
	ret.Inst.PublicIpv4Addr = d.Get("public_ipv4_addr").(string)
	ret.Inst.PublicIpv6Addr = d.Get("public_ipv6_addr").(string)
	ret.Inst.SpoofDetectCfg = getObjectDdosTemplateUdpSpoofDetectCfg(d.Get("spoof_detect_cfg").([]interface{}))
	ret.Inst.TokenAuthentication = d.Get("token_authentication").(int)
	ret.Inst.TokenAuthenticationFormula = d.Get("token_authentication_formula").(string)
	ret.Inst.TokenAuthenticationHwAssistDisable = d.Get("token_authentication_hw_assist_disable").(int)
	ret.Inst.TokenAuthenticationPublicAddress = d.Get("token_authentication_public_address").(int)
	ret.Inst.TokenAuthenticationSaltPrefix = d.Get("token_authentication_salt_prefix").(int)
	ret.Inst.TokenAuthenticationSaltPrefixCurr = d.Get("token_authentication_salt_prefix_curr").(int)
	ret.Inst.TokenAuthenticationSaltPrefixPrev = d.Get("token_authentication_salt_prefix_prev").(int)
	ret.Inst.TunnelEncap = getObjectDdosTemplateUdpTunnelEncap(d.Get("tunnel_encap").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
