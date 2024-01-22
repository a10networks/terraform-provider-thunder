package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetectionOutboundDetectionIndicator() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_detection_outbound_detection_indicator`: Outbound indicator configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneDetectionOutboundDetectionIndicatorCreate,
		UpdateContext: resourceDdosDstZoneDetectionOutboundDetectionIndicatorUpdate,
		ReadContext:   resourceDdosDstZoneDetectionOutboundDetectionIndicatorRead,
		DeleteContext: resourceDdosDstZoneDetectionOutboundDetectionIndicatorDelete,

		Schema: map[string]*schema.Schema{
			"data_packet_size": {
				Type: schema.TypeInt, Optional: true, Description: "Expected minimal data size",
			},
			"tcp_window_size": {
				Type: schema.TypeInt, Optional: true, Description: "Expected minimal window size",
			},
			"threshold_large_num": {
				Type: schema.TypeInt, Optional: true, Description: "Threshold for each geo-location",
			},
			"threshold_num": {
				Type: schema.TypeInt, Optional: true, Description: "Threshold for each geo-location",
			},
			"threshold_str": {
				Type: schema.TypeString, Optional: true, Description: "Threshold for each geo-location (Non-zero floating point)",
			},
			"type": {
				Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'pkt-drop-rate': rate of packets got dropped; 'bit-rate': rate of incoming bits; 'pkt-drop-ratio': ratio of incoming packet rate divided by the rate of dropping packets; 'bytes-to-bytes-from-ratio': ratio of incoming packet rate divided by the rate of outgoing packets; 'syn-rate': rate on incoming SYN packets; 'fin-rate': rate on incoming FIN packets; 'rst-rate': rate of incoming RST packets; 'small-window-ack-rate': rate of small window advertisement; 'empty-ack-rate': rate of incoming packets which have no payload; 'small-payload-rate': rate of short payload packet; 'syn-fin-ratio': ratio of incoming SYN packet rate divided by the rate of incoming FIN packets;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_name": {
				Type: schema.TypeString, Required: true, Description: "ZoneName",
			},
		},
	}
}
func resourceDdosDstZoneDetectionOutboundDetectionIndicatorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionOutboundDetectionIndicatorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionOutboundDetectionIndicator(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionOutboundDetectionIndicatorRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneDetectionOutboundDetectionIndicatorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionOutboundDetectionIndicatorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionOutboundDetectionIndicator(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionOutboundDetectionIndicatorRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneDetectionOutboundDetectionIndicatorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionOutboundDetectionIndicatorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionOutboundDetectionIndicator(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneDetectionOutboundDetectionIndicatorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionOutboundDetectionIndicatorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionOutboundDetectionIndicator(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneDetectionOutboundDetectionIndicator(d *schema.ResourceData) edpt.DdosDstZoneDetectionOutboundDetectionIndicator {
	var ret edpt.DdosDstZoneDetectionOutboundDetectionIndicator
	ret.Inst.DataPacketSize = d.Get("data_packet_size").(int)
	ret.Inst.TcpWindowSize = d.Get("tcp_window_size").(int)
	ret.Inst.ThresholdLargeNum = d.Get("threshold_large_num").(int)
	ret.Inst.ThresholdNum = d.Get("threshold_num").(int)
	ret.Inst.ThresholdStr = d.Get("threshold_str").(string)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
