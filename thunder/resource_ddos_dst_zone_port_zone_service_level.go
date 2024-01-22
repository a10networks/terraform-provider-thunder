package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceLevel() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_zone_service_level`: Policy Level Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortZoneServiceLevelCreate,
		UpdateContext: resourceDdosDstZonePortZoneServiceLevelUpdate,
		ReadContext:   resourceDdosDstZonePortZoneServiceLevelRead,
		DeleteContext: resourceDdosDstZonePortZoneServiceLevelDelete,

		Schema: map[string]*schema.Schema{
			"apply_extracted_filters": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Apply extracted filters from this level",
			},
			"close_sessions_for_unauth_sources": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Close session for unauthenticated sources",
			},
			"glid_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
			},
			"indicator_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'concurrent-conns': number of concurrent connections; 'conn-miss-rate': rate of incoming packets for which no previously established connection exists; 'syn-rate': rate on incoming SYN packets; 'fin-rate': rate on incoming FIN packets; 'rst-rate': rate of incoming RST packets; 'small-window-ack-rate': rate of small window advertisement; 'empty-ack-rate': rate of incoming packets which have no payload; 'small-payload-rate': rate of short payload packet; 'syn-fin-ratio': ratio of incoming SYN packet rate divided by the rate of incoming FIN packets; 'cpu-utilization': average data CPU utilization; 'interface-utilization': outside interface utilization;",
						},
						"tcp_window_size": {
							Type: schema.TypeInt, Optional: true, Description: "Expected minimal window size",
						},
						"data_packet_size": {
							Type: schema.TypeInt, Optional: true, Description: "Expected minimal data size",
						},
						"score": {
							Type: schema.TypeInt, Optional: true, Description: "Score corresponding to the indicator",
						},
						"src_threshold_num": {
							Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
						},
						"src_threshold_large_num": {
							Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
						},
						"src_threshold_str": {
							Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
						},
						"src_violation_actions": {
							Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this src indicator threshold reaches",
						},
						"zone_threshold_large_num": {
							Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
						},
						"zone_threshold_num": {
							Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
						},
						"zone_threshold_str": {
							Type: schema.TypeString, Optional: true, Description: "Threshold for the entire zone (Non-zero floating point)",
						},
						"zone_violation_actions": {
							Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this zone indicator threshold reaches",
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
			"level_num": {
				Type: schema.TypeString, Required: true, Description: "'0': Default policy level; '1': Policy level 1; '2': Policy level 2; '3': Policy level 3; '4': Policy level 4;",
			},
			"src_default_glid": {
				Type: schema.TypeString, Optional: true, Description: "Global limit ID",
			},
			"src_escalation_score": {
				Type: schema.TypeInt, Optional: true, Description: "Source activation score of this level",
			},
			"src_violation_actions": {
				Type: schema.TypeString, Optional: true, Description: "Violation actions apply due to source escalate from this level",
			},
			"start_pattern_recognition": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Start pattern recognition from this level",
			},
			"start_signature_extraction": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Start signature extraction from this level",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_escalation_score": {
				Type: schema.TypeInt, Optional: true, Description: "Zone activation score of this level",
			},
			"zone_template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"quic": {
							Type: schema.TypeString, Optional: true, Description: "DDOS quic template",
						},
						"dns": {
							Type: schema.TypeString, Optional: true, Description: "DDOS dns template",
						},
						"http": {
							Type: schema.TypeString, Optional: true, Description: "DDOS http template",
						},
						"ssl_l4": {
							Type: schema.TypeString, Optional: true, Description: "DDOS ssl-l4 template",
						},
						"sip": {
							Type: schema.TypeString, Optional: true, Description: "DDOS sip template",
						},
						"tcp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS tcp template",
						},
						"udp": {
							Type: schema.TypeString, Optional: true, Description: "DDOS udp template",
						},
						"encap": {
							Type: schema.TypeString, Optional: true, Description: "DDOS encap template (IPv6-over-IPv4 / IPv4-over-IPv6 are not supported.)",
						},
					},
				},
			},
			"zone_violation_actions": {
				Type: schema.TypeString, Optional: true, Description: "Violation actions apply due to zone escalate from this level",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}
func resourceDdosDstZonePortZoneServiceLevelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceLevelCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceLevel(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceLevelRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceLevelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceLevelUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceLevel(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceLevelRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortZoneServiceLevelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceLevelDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceLevel(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceLevelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceLevelRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceLevel(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZonePortZoneServiceLevelIndicatorList(d []interface{}) []edpt.DdosDstZonePortZoneServiceLevelIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZonePortZoneServiceLevelIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZonePortZoneServiceLevelIndicatorList
		oi.Type = in["type"].(string)
		oi.TcpWindowSize = in["tcp_window_size"].(int)
		oi.DataPacketSize = in["data_packet_size"].(int)
		oi.Score = in["score"].(int)
		oi.SrcThresholdNum = in["src_threshold_num"].(int)
		oi.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		oi.SrcThresholdStr = in["src_threshold_str"].(string)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdStr = in["zone_threshold_str"].(string)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZonePortZoneServiceLevelZoneTemplate(d []interface{}) edpt.DdosDstZonePortZoneServiceLevelZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZonePortZoneServiceLevelZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Quic = in["quic"].(string)
		ret.Dns = in["dns"].(string)
		ret.Http = in["http"].(string)
		ret.SslL4 = in["ssl_l4"].(string)
		ret.Sip = in["sip"].(string)
		ret.Tcp = in["tcp"].(string)
		ret.Udp = in["udp"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZonePortZoneServiceLevel(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceLevel {
	var ret edpt.DdosDstZonePortZoneServiceLevel
	ret.Inst.ApplyExtractedFilters = d.Get("apply_extracted_filters").(int)
	ret.Inst.CloseSessionsForUnauthSources = d.Get("close_sessions_for_unauth_sources").(int)
	ret.Inst.GlidAction = d.Get("glid_action").(string)
	ret.Inst.IndicatorList = getSliceDdosDstZonePortZoneServiceLevelIndicatorList(d.Get("indicator_list").([]interface{}))
	ret.Inst.LevelNum = d.Get("level_num").(string)
	ret.Inst.SrcDefaultGlid = d.Get("src_default_glid").(string)
	ret.Inst.SrcEscalationScore = d.Get("src_escalation_score").(int)
	ret.Inst.SrcViolationActions = d.Get("src_violation_actions").(string)
	ret.Inst.StartPatternRecognition = d.Get("start_pattern_recognition").(int)
	ret.Inst.StartSignatureExtraction = d.Get("start_signature_extraction").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneEscalationScore = d.Get("zone_escalation_score").(int)
	ret.Inst.ZoneTemplate = getObjectDdosDstZonePortZoneServiceLevelZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneViolationActions = d.Get("zone_violation_actions").(string)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.PortNum = d.Get("port_num").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	return ret
}
