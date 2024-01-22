package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneIpProtoProtoNumberLevel() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_ip_proto_proto_number_level`: Policy Level Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneIpProtoProtoNumberLevelCreate,
		UpdateContext: resourceDdosDstZoneIpProtoProtoNumberLevelUpdate,
		ReadContext:   resourceDdosDstZoneIpProtoProtoNumberLevelRead,
		DeleteContext: resourceDdosDstZoneIpProtoProtoNumberLevelDelete,

		Schema: map[string]*schema.Schema{
			"glid_action": {
				Type: schema.TypeString, Optional: true, Description: "'drop': Drop packets for glid exceed (Default); 'blacklist-src': Blacklist-src for glid exceed; 'ignore': Do nothing for glid exceed;",
			},
			"indicator_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'frag-rate': rate of incoming fragmented packets; 'cpu-utilization': average data CPU utilization; 'interface-utilization': outside interface utilization;",
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
						"zone_threshold_num": {
							Type: schema.TypeInt, Optional: true, Description: "Threshold for the entire zone",
						},
						"zone_threshold_large_num": {
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
						"ip_proto": {
							Type: schema.TypeString, Optional: true, Description: "DDOS ip-proto template",
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
			"protocol_num": {
				Type: schema.TypeString, Required: true, Description: "ProtocolNum",
			},
		},
	}
}
func resourceDdosDstZoneIpProtoProtoNumberLevelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberLevelCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberLevel(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNumberLevelRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNumberLevelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberLevelUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberLevel(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneIpProtoProtoNumberLevelRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneIpProtoProtoNumberLevelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberLevelDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberLevel(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneIpProtoProtoNumberLevelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneIpProtoProtoNumberLevelRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneIpProtoProtoNumberLevel(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZoneIpProtoProtoNumberLevelIndicatorList(d []interface{}) []edpt.DdosDstZoneIpProtoProtoNumberLevelIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneIpProtoProtoNumberLevelIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneIpProtoProtoNumberLevelIndicatorList
		oi.Type = in["type"].(string)
		oi.DataPacketSize = in["data_packet_size"].(int)
		oi.Score = in["score"].(int)
		oi.SrcThresholdNum = in["src_threshold_num"].(int)
		oi.SrcThresholdLargeNum = in["src_threshold_large_num"].(int)
		oi.SrcThresholdStr = in["src_threshold_str"].(string)
		oi.SrcViolationActions = in["src_violation_actions"].(string)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		oi.ZoneThresholdStr = in["zone_threshold_str"].(string)
		oi.ZoneViolationActions = in["zone_violation_actions"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosDstZoneIpProtoProtoNumberLevelZoneTemplate(d []interface{}) edpt.DdosDstZoneIpProtoProtoNumberLevelZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosDstZoneIpProtoProtoNumberLevelZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.IpProto = in["ip_proto"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func dataToEndpointDdosDstZoneIpProtoProtoNumberLevel(d *schema.ResourceData) edpt.DdosDstZoneIpProtoProtoNumberLevel {
	var ret edpt.DdosDstZoneIpProtoProtoNumberLevel
	ret.Inst.GlidAction = d.Get("glid_action").(string)
	ret.Inst.IndicatorList = getSliceDdosDstZoneIpProtoProtoNumberLevelIndicatorList(d.Get("indicator_list").([]interface{}))
	ret.Inst.LevelNum = d.Get("level_num").(string)
	ret.Inst.SrcDefaultGlid = d.Get("src_default_glid").(string)
	ret.Inst.SrcEscalationScore = d.Get("src_escalation_score").(int)
	ret.Inst.SrcViolationActions = d.Get("src_violation_actions").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneEscalationScore = d.Get("zone_escalation_score").(int)
	ret.Inst.ZoneTemplate = getObjectDdosDstZoneIpProtoProtoNumberLevelZoneTemplate(d.Get("zone_template").([]interface{}))
	ret.Inst.ZoneViolationActions = d.Get("zone_violation_actions").(string)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.ProtocolNum = d.Get("protocol_num").(string)
	return ret
}
