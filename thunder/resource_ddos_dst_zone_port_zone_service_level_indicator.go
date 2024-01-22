package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZonePortZoneServiceLevelIndicator() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_port_zone_service_level_indicator`: DDoS Indicator Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZonePortZoneServiceLevelIndicatorCreate,
		UpdateContext: resourceDdosDstZonePortZoneServiceLevelIndicatorUpdate,
		ReadContext:   resourceDdosDstZonePortZoneServiceLevelIndicatorRead,
		DeleteContext: resourceDdosDstZonePortZoneServiceLevelIndicatorDelete,

		Schema: map[string]*schema.Schema{
			"data_packet_size": {
				Type: schema.TypeInt, Optional: true, Description: "Expected minimal data size",
			},
			"score": {
				Type: schema.TypeInt, Optional: true, Description: "Score corresponding to the indicator",
			},
			"src_threshold_large_num": {
				Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
			},
			"src_threshold_num": {
				Type: schema.TypeInt, Optional: true, Description: "Indicator per-src threshold",
			},
			"src_threshold_str": {
				Type: schema.TypeString, Optional: true, Description: "Indicator per-src threshold (Non-zero floating point)",
			},
			"src_violation_actions": {
				Type: schema.TypeString, Optional: true, Description: "Violation actions to use when this src indicator threshold reaches",
			},
			"tcp_window_size": {
				Type: schema.TypeInt, Optional: true, Description: "Expected minimal window size",
			},
			"type": {
				Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'concurrent-conns': number of concurrent connections; 'conn-miss-rate': rate of incoming packets for which no previously established connection exists; 'syn-rate': rate on incoming SYN packets; 'fin-rate': rate on incoming FIN packets; 'rst-rate': rate of incoming RST packets; 'small-window-ack-rate': rate of small window advertisement; 'empty-ack-rate': rate of incoming packets which have no payload; 'small-payload-rate': rate of short payload packet; 'syn-fin-ratio': ratio of incoming SYN packet rate divided by the rate of incoming FIN packets; 'cpu-utilization': average data CPU utilization; 'interface-utilization': outside interface utilization;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
			"level_num": {
				Type: schema.TypeString, Required: true, Description: "LevelNum",
			},
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
		},
	}
}
func resourceDdosDstZonePortZoneServiceLevelIndicatorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceLevelIndicatorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceLevelIndicator(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceLevelIndicatorRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceLevelIndicatorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceLevelIndicatorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceLevelIndicator(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZonePortZoneServiceLevelIndicatorRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZonePortZoneServiceLevelIndicatorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceLevelIndicatorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceLevelIndicator(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZonePortZoneServiceLevelIndicatorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZonePortZoneServiceLevelIndicatorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZonePortZoneServiceLevelIndicator(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZonePortZoneServiceLevelIndicator(d *schema.ResourceData) edpt.DdosDstZonePortZoneServiceLevelIndicator {
	var ret edpt.DdosDstZonePortZoneServiceLevelIndicator
	ret.Inst.DataPacketSize = d.Get("data_packet_size").(int)
	ret.Inst.Score = d.Get("score").(int)
	ret.Inst.SrcThresholdLargeNum = d.Get("src_threshold_large_num").(int)
	ret.Inst.SrcThresholdNum = d.Get("src_threshold_num").(int)
	ret.Inst.SrcThresholdStr = d.Get("src_threshold_str").(string)
	ret.Inst.SrcViolationActions = d.Get("src_violation_actions").(string)
	ret.Inst.TcpWindowSize = d.Get("tcp_window_size").(int)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneThresholdLargeNum = d.Get("zone_threshold_large_num").(int)
	ret.Inst.ZoneThresholdNum = d.Get("zone_threshold_num").(int)
	ret.Inst.ZoneThresholdStr = d.Get("zone_threshold_str").(string)
	ret.Inst.ZoneViolationActions = d.Get("zone_violation_actions").(string)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.LevelNum = d.Get("level_num").(string)
	ret.Inst.PortNum = d.Get("port_num").(string)
	return ret
}
