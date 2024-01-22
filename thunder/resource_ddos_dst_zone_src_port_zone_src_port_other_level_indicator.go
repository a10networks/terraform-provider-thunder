package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicator() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_src_port_zone_src_port_other_level_indicator`: DDoS Indicator Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorCreate,
		UpdateContext: resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorUpdate,
		ReadContext:   resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorRead,
		DeleteContext: resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'bit-rate': rate of incoming bits;",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_threshold_large_num": {
				Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the src-port",
			},
			"zone_threshold_num": {
				Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the src-port",
			},
			"port_other": {
				Type: schema.TypeString, Required: true, Description: "PortOther",
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
		},
	}
}
func resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicator(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicator(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicator(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicatorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicator(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneSrcPortZoneSrcPortOtherLevelIndicator(d *schema.ResourceData) edpt.DdosDstZoneSrcPortZoneSrcPortOtherLevelIndicator {
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortOtherLevelIndicator
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneThresholdLargeNum = d.Get("zone_threshold_large_num").(int)
	ret.Inst.ZoneThresholdNum = d.Get("zone_threshold_num").(int)
	ret.Inst.PortOther = d.Get("port_other").(string)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.LevelNum = d.Get("level_num").(string)
	return ret
}
