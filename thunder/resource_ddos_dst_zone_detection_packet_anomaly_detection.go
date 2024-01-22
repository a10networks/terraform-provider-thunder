package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetectionPacketAnomalyDetection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_detection_packet_anomaly_detection`: DDos Packet Anomaly detections\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneDetectionPacketAnomalyDetectionCreate,
		UpdateContext: resourceDdosDstZoneDetectionPacketAnomalyDetectionUpdate,
		ReadContext:   resourceDdosDstZoneDetectionPacketAnomalyDetectionRead,
		DeleteContext: resourceDdosDstZoneDetectionPacketAnomalyDetectionDelete,

		Schema: map[string]*schema.Schema{
			"configuration": {
				Type: schema.TypeString, Required: true, Description: "'configuration': configuration;",
			},
			"indicator_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Required: true, Description: "'port-zero-pkt-rate': Port Zero Packet Rate (default 100 packet per second);",
						},
						"threshold_num": {
							Type: schema.TypeInt, Optional: true, Default: 100, Description: "Threshold for each indicator",
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
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable packet anomaly; 'disable': Disable packet anomaly;",
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
func resourceDdosDstZoneDetectionPacketAnomalyDetectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionPacketAnomalyDetectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionPacketAnomalyDetection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionPacketAnomalyDetectionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneDetectionPacketAnomalyDetectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionPacketAnomalyDetectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionPacketAnomalyDetection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionPacketAnomalyDetectionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneDetectionPacketAnomalyDetectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionPacketAnomalyDetectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionPacketAnomalyDetection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneDetectionPacketAnomalyDetectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionPacketAnomalyDetectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionPacketAnomalyDetection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorList(d []interface{}) []edpt.DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionPacketAnomalyDetectionIndicatorList
		oi.Type = in["type"].(string)
		oi.ThresholdNum = in["threshold_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneDetectionPacketAnomalyDetection(d *schema.ResourceData) edpt.DdosDstZoneDetectionPacketAnomalyDetection {
	var ret edpt.DdosDstZoneDetectionPacketAnomalyDetection
	ret.Inst.Configuration = d.Get("configuration").(string)
	ret.Inst.IndicatorList = getSliceDdosDstZoneDetectionPacketAnomalyDetectionIndicatorList(d.Get("indicator_list").([]interface{}))
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
