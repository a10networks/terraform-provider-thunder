package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneDetectionVictimIpDetectionIndicator() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_detection_victim_ip_detection_indicator`: DDos Host Detections configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneDetectionVictimIpDetectionIndicatorCreate,
		UpdateContext: resourceDdosDstZoneDetectionVictimIpDetectionIndicatorUpdate,
		ReadContext:   resourceDdosDstZoneDetectionVictimIpDetectionIndicatorRead,
		DeleteContext: resourceDdosDstZoneDetectionVictimIpDetectionIndicatorDelete,

		Schema: map[string]*schema.Schema{
			"ip_threshold_num": {
				Type: schema.TypeInt, Optional: true, Description: "Threshold for IP",
			},
			"type": {
				Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'reverse-pkt-rate': rate of reverse coming packets; 'fwd-byte-rate': rate of incoming bytes; 'rev-byte-rate': rate of reverse coming bytes;",
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
func resourceDdosDstZoneDetectionVictimIpDetectionIndicatorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionVictimIpDetectionIndicatorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionVictimIpDetectionIndicator(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionVictimIpDetectionIndicatorRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneDetectionVictimIpDetectionIndicatorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionVictimIpDetectionIndicatorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionVictimIpDetectionIndicator(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneDetectionVictimIpDetectionIndicatorRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneDetectionVictimIpDetectionIndicatorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionVictimIpDetectionIndicatorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionVictimIpDetectionIndicator(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneDetectionVictimIpDetectionIndicatorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneDetectionVictimIpDetectionIndicatorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneDetectionVictimIpDetectionIndicator(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneDetectionVictimIpDetectionIndicator(d *schema.ResourceData) edpt.DdosDstZoneDetectionVictimIpDetectionIndicator {
	var ret edpt.DdosDstZoneDetectionVictimIpDetectionIndicator
	ret.Inst.IpThresholdNum = d.Get("ip_threshold_num").(int)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
