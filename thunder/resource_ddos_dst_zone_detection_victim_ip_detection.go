package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetectionVictimIpDetection() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_detection_victim_ip_detection`: DDos Host Detections configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneDetectionVictimIpDetectionCreate,
		UpdateContext: resourceDdosDstZoneDetectionVictimIpDetectionUpdate,
		ReadContext:   resourceDdosDstZoneDetectionVictimIpDetectionRead,
		DeleteContext: resourceDdosDstZoneDetectionVictimIpDetectionDelete,

		Schema: map[string]*schema.Schema{
			"configuration": {
				Type: schema.TypeString, Required: true, Description: "'configuration': configuration;",
			},
			"histogram_toggle": {
				Type: schema.TypeString, Optional: true, Default: "histogram-disable", Description: "'histogram-enable': Enable histogram statistics of victim IP detection; 'histogram-disable': Disable histogram statistics of victim IP detection;",
			},
			"indicator_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'reverse-pkt-rate': rate of reverse coming packets; 'fwd-byte-rate': rate of incoming bytes; 'rev-byte-rate': rate of reverse coming bytes;",
						},
						"ip_threshold_num": {
							Type: schema.TypeInt, Optional: true, Description: "Threshold for IP",
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
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'enable': Enable victim IP detection; 'disable': Disable victim IP detection;",
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
func resourceDdosDstZoneDetectionVictimIpDetectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionVictimIpDetectionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionVictimIpDetection(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionVictimIpDetectionRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneDetectionVictimIpDetectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionVictimIpDetectionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionVictimIpDetection(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionVictimIpDetectionRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneDetectionVictimIpDetectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionVictimIpDetectionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionVictimIpDetection(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneDetectionVictimIpDetectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionVictimIpDetectionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionVictimIpDetection(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZoneDetectionVictimIpDetectionIndicatorList(d []interface{}) []edpt.DdosDstZoneDetectionVictimIpDetectionIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneDetectionVictimIpDetectionIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneDetectionVictimIpDetectionIndicatorList
		oi.Type = in["type"].(string)
		oi.IpThresholdNum = in["ip_threshold_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneDetectionVictimIpDetection(d *schema.ResourceData) edpt.DdosDstZoneDetectionVictimIpDetection {
	var ret edpt.DdosDstZoneDetectionVictimIpDetection
	ret.Inst.Configuration = d.Get("configuration").(string)
	ret.Inst.HistogramToggle = d.Get("histogram_toggle").(string)
	ret.Inst.IndicatorList = getSliceDdosDstZoneDetectionVictimIpDetectionIndicatorList(d.Get("indicator_list").([]interface{}))
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
