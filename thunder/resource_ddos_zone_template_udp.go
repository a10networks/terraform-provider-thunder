package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosZoneTemplateUdp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_zone_template_udp`: UDP template configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosZoneTemplateUdpCreate,
		UpdateContext: resourceDdosZoneTemplateUdpUpdate,
		ReadContext:   resourceDdosZoneTemplateUdpRead,
		DeleteContext: resourceDdosZoneTemplateUdpDelete,

		Schema: map[string]*schema.Schema{
			"age": {
				Type: schema.TypeInt, Optional: true, Default: 2, Description: "Configure session age(in minutes) for UDP sessions",
			},
			"filter_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"udp_filter_name": {
							Type: schema.TypeString, Required: true, Description: "",
						},
						"udp_filter_seq": {
							Type: schema.TypeInt, Optional: true, Description: "Sequence number",
						},
						"udp_filter_regex": {
							Type: schema.TypeString, Optional: true, Description: "Regex Expression",
						},
						"udp_filter_inverse_match": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inverse the result of the matching",
						},
						"byte_offset_filter": {
							Type: schema.TypeString, Optional: true, Description: "Filter using Berkeley Packet Filter syntax",
						},
						"udp_filter_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take",
						},
						"udp_filter_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets (Default); 'ignore': Take no action; 'blacklist-src': Blacklist-src; 'authenticate-src': Authenticate-src;",
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
			"filter_match_type": {
				Type: schema.TypeString, Optional: true, Default: "default", Description: "'default': Stop matching on drop/blacklist action; 'stop-on-first-match': Stop matching on first match;",
			},
			"known_resp_src_port_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"known_resp_src_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Take action if src-port is less than 1024",
						},
						"known_resp_src_port_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for well-known src-port",
						},
						"known_resp_src_port_action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets from well-known src-port(Default); 'blacklist-src': Blacklist-src from well-known src-port; 'ignore': Ignore well-known src-port;",
						},
						"exclude_src_resp_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Exclude src port equal to dst port",
						},
					},
				},
			},
			"max_payload_size_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_payload_size": {
							Type: schema.TypeInt, Optional: true, Description: "Maximum UDP payload size for each single packet",
						},
						"max_payload_size_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for max-payload-size exceed",
						},
						"max_payload_size_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for max-payload-size exceed (Default); 'blacklist-src': Blacklist-src for max-payload-size exceed; 'ignore': Do nothing for max-payload-size exceed;",
						},
					},
				},
			},
			"min_payload_size_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_payload_size": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum UDP payload size for each single packet",
						},
						"min_payload_size_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for min-payload-size exceed",
						},
						"min_payload_size_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for min-payload-size (Default); 'blacklist-src': Blacklist-src for min-payload-size; 'ignore': Do nothing for min-payload-size exceed;",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DDOS UDP Template Name",
			},
			"ntp_monlist_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ntp_monlist": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Take action for ntp monlist request/response",
						},
						"ntp_monlist_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for ntp-monlist",
						},
						"ntp_monlist_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for ntp-monlist (Default); 'blacklist-src': Blacklist-src for ntp-monlist; 'ignore': Ignore ntp-monlist;",
						},
					},
				},
			},
			"per_conn_pkt_rate_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"per_conn_pkt_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Packet rate limit per connection per rate-interval",
						},
						"per_conn_pkt_rate_action_list_name": {
							Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for per-conn-pkt-rate exceed",
						},
						"per_conn_pkt_rate_action": {
							Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': Drop packets for per-conn-pkt-rate exceed (Default); 'blacklist-src': help Blacklist-src for per-conn-pkt-rate exceed; 'ignore': Ignore per-conn-pkt-rate-exceed;",
						},
					},
				},
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
			"spoof_detect_fail_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets (Default); 'blacklist-src': Blacklist-src for spoof-detect fail;",
			},
			"spoof_detect_fail_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for failing the authentication",
			},
			"spoof_detect_min_delay": {
				Type: schema.TypeInt, Optional: true, Description: "Optional minimum delay between UDP retransmits for authentication to pass, unit is specified by min-delay-interval",
			},
			"spoof_detect_min_delay_interval": {
				Type: schema.TypeString, Optional: true, Default: "1sec", Description: "'100ms': 100ms; '1sec': 1sec;",
			},
			"spoof_detect_pass_action": {
				Type: schema.TypeString, Optional: true, Description: "'authenticate-src': authenticate-src (Default);",
			},
			"spoof_detect_pass_action_list_name": {
				Type: schema.TypeString, Optional: true, Description: "Configure action-list to take for passing the authentication",
			},
			"spoof_detect_retry_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Timeout in seconds",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosZoneTemplateUdpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateUdpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateUdp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateUdpRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosZoneTemplateUdpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateUdpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateUdp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosZoneTemplateUdpRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosZoneTemplateUdpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateUdpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateUdp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosZoneTemplateUdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosZoneTemplateUdpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosZoneTemplateUdp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosZoneTemplateUdpFilterList(d []interface{}) []edpt.DdosZoneTemplateUdpFilterList {

	count1 := len(d)
	ret := make([]edpt.DdosZoneTemplateUdpFilterList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosZoneTemplateUdpFilterList
		oi.UdpFilterName = in["udp_filter_name"].(string)
		oi.UdpFilterSeq = in["udp_filter_seq"].(int)
		oi.UdpFilterRegex = in["udp_filter_regex"].(string)
		oi.UdpFilterInverseMatch = in["udp_filter_inverse_match"].(int)
		oi.ByteOffsetFilter = in["byte_offset_filter"].(string)
		oi.UdpFilterActionListName = in["udp_filter_action_list_name"].(string)
		oi.UdpFilterAction = in["udp_filter_action"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosZoneTemplateUdpKnownRespSrcPortCfg(d []interface{}) edpt.DdosZoneTemplateUdpKnownRespSrcPortCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateUdpKnownRespSrcPortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KnownRespSrcPort = in["known_resp_src_port"].(int)
		ret.KnownRespSrcPortActionListName = in["known_resp_src_port_action_list_name"].(string)
		ret.KnownRespSrcPortAction = in["known_resp_src_port_action"].(string)
		ret.ExcludeSrcRespPort = in["exclude_src_resp_port"].(int)
	}
	return ret
}

func getObjectDdosZoneTemplateUdpMaxPayloadSizeCfg(d []interface{}) edpt.DdosZoneTemplateUdpMaxPayloadSizeCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateUdpMaxPayloadSizeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MaxPayloadSize = in["max_payload_size"].(int)
		ret.MaxPayloadSizeActionListName = in["max_payload_size_action_list_name"].(string)
		ret.MaxPayloadSizeAction = in["max_payload_size_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateUdpMinPayloadSizeCfg(d []interface{}) edpt.DdosZoneTemplateUdpMinPayloadSizeCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateUdpMinPayloadSizeCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MinPayloadSize = in["min_payload_size"].(int)
		ret.MinPayloadSizeActionListName = in["min_payload_size_action_list_name"].(string)
		ret.MinPayloadSizeAction = in["min_payload_size_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateUdpNtpMonlistCfg(d []interface{}) edpt.DdosZoneTemplateUdpNtpMonlistCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateUdpNtpMonlistCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.NtpMonlist = in["ntp_monlist"].(int)
		ret.NtpMonlistActionListName = in["ntp_monlist_action_list_name"].(string)
		ret.NtpMonlistAction = in["ntp_monlist_action"].(string)
	}
	return ret
}

func getObjectDdosZoneTemplateUdpPerConnPktRateCfg(d []interface{}) edpt.DdosZoneTemplateUdpPerConnPktRateCfg {

	count1 := len(d)
	var ret edpt.DdosZoneTemplateUdpPerConnPktRateCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PerConnPktRateLimit = in["per_conn_pkt_rate_limit"].(int)
		ret.PerConnPktRateActionListName = in["per_conn_pkt_rate_action_list_name"].(string)
		ret.PerConnPktRateAction = in["per_conn_pkt_rate_action"].(string)
	}
	return ret
}

func dataToEndpointDdosZoneTemplateUdp(d *schema.ResourceData) edpt.DdosZoneTemplateUdp {
	var ret edpt.DdosZoneTemplateUdp
	ret.Inst.Age = d.Get("age").(int)
	ret.Inst.FilterList = getSliceDdosZoneTemplateUdpFilterList(d.Get("filter_list").([]interface{}))
	ret.Inst.FilterMatchType = d.Get("filter_match_type").(string)
	ret.Inst.KnownRespSrcPortCfg = getObjectDdosZoneTemplateUdpKnownRespSrcPortCfg(d.Get("known_resp_src_port_cfg").([]interface{}))
	ret.Inst.MaxPayloadSizeCfg = getObjectDdosZoneTemplateUdpMaxPayloadSizeCfg(d.Get("max_payload_size_cfg").([]interface{}))
	ret.Inst.MinPayloadSizeCfg = getObjectDdosZoneTemplateUdpMinPayloadSizeCfg(d.Get("min_payload_size_cfg").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.NtpMonlistCfg = getObjectDdosZoneTemplateUdpNtpMonlistCfg(d.Get("ntp_monlist_cfg").([]interface{}))
	ret.Inst.PerConnPktRateCfg = getObjectDdosZoneTemplateUdpPerConnPktRateCfg(d.Get("per_conn_pkt_rate_cfg").([]interface{}))
	ret.Inst.PerConnRateInterval = d.Get("per_conn_rate_interval").(string)
	ret.Inst.PreviousSaltTimeout = d.Get("previous_salt_timeout").(int)
	ret.Inst.PublicIpv4Addr = d.Get("public_ipv4_addr").(string)
	ret.Inst.PublicIpv6Addr = d.Get("public_ipv6_addr").(string)
	ret.Inst.SpoofDetectFailAction = d.Get("spoof_detect_fail_action").(string)
	ret.Inst.SpoofDetectFailActionListName = d.Get("spoof_detect_fail_action_list_name").(string)
	ret.Inst.SpoofDetectMinDelay = d.Get("spoof_detect_min_delay").(int)
	ret.Inst.SpoofDetectMinDelayInterval = d.Get("spoof_detect_min_delay_interval").(string)
	ret.Inst.SpoofDetectPassAction = d.Get("spoof_detect_pass_action").(string)
	ret.Inst.SpoofDetectPassActionListName = d.Get("spoof_detect_pass_action_list_name").(string)
	ret.Inst.SpoofDetectRetryTimeout = d.Get("spoof_detect_retry_timeout").(int)
	ret.Inst.TokenAuthentication = d.Get("token_authentication").(int)
	ret.Inst.TokenAuthenticationFormula = d.Get("token_authentication_formula").(string)
	ret.Inst.TokenAuthenticationHwAssistDisable = d.Get("token_authentication_hw_assist_disable").(int)
	ret.Inst.TokenAuthenticationPublicAddress = d.Get("token_authentication_public_address").(int)
	ret.Inst.TokenAuthenticationSaltPrefix = d.Get("token_authentication_salt_prefix").(int)
	ret.Inst.TokenAuthenticationSaltPrefixCurr = d.Get("token_authentication_salt_prefix_curr").(int)
	ret.Inst.TokenAuthenticationSaltPrefixPrev = d.Get("token_authentication_salt_prefix_prev").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
