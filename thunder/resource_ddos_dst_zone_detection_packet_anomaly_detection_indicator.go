package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicator() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_detection_packet_anomaly_detection_indicator`: Packet Anomaly indicator configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorCreate,
		UpdateContext: resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorUpdate,
		ReadContext:   resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorRead,
		DeleteContext: resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorDelete,

		Schema: map[string]*schema.Schema{
			"threshold_num": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Threshold for each indicator",
			},
			"type": {
				Type: schema.TypeString, Required: true, Description: "'port-zero-pkt-rate': Port Zero Packet Rate (default 100 packet per second);",
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
func resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionPacketAnomalyDetectionIndicator(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionPacketAnomalyDetectionIndicator(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionPacketAnomalyDetectionIndicator(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionPacketAnomalyDetectionIndicator(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneDetectionPacketAnomalyDetectionIndicator(d *schema.ResourceData) edpt.DdosDstZoneDetectionPacketAnomalyDetectionIndicator {
	var ret edpt.DdosDstZoneDetectionPacketAnomalyDetectionIndicator
	ret.Inst.ThresholdNum = d.Get("threshold_num").(int)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
