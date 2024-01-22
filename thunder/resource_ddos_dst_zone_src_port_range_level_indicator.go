package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcPortRangeLevelIndicator() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_src_port_range_level_indicator`: DDoS Indicator Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneSrcPortRangeLevelIndicatorCreate,
		UpdateContext: resourceDdosDstZoneSrcPortRangeLevelIndicatorUpdate,
		ReadContext:   resourceDdosDstZoneSrcPortRangeLevelIndicatorRead,
		DeleteContext: resourceDdosDstZoneSrcPortRangeLevelIndicatorDelete,

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
				Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the port-range",
			},
			"zone_threshold_num": {
				Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the port-range",
			},
			"src_port_range_end": {
				Type: schema.TypeString, Required: true, Description: "SrcPortRangeEnd",
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
			"src_port_range_start": {
				Type: schema.TypeString, Required: true, Description: "SrcPortRangeStart",
			},
		},
	}
}
func resourceDdosDstZoneSrcPortRangeLevelIndicatorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortRangeLevelIndicatorCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortRangeLevelIndicator(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcPortRangeLevelIndicatorRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneSrcPortRangeLevelIndicatorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortRangeLevelIndicatorUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortRangeLevelIndicator(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcPortRangeLevelIndicatorRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneSrcPortRangeLevelIndicatorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortRangeLevelIndicatorDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortRangeLevelIndicator(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneSrcPortRangeLevelIndicatorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortRangeLevelIndicatorRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortRangeLevelIndicator(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneSrcPortRangeLevelIndicator(d *schema.ResourceData) edpt.DdosDstZoneSrcPortRangeLevelIndicator {
	var ret edpt.DdosDstZoneSrcPortRangeLevelIndicator
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneThresholdLargeNum = d.Get("zone_threshold_large_num").(int)
	ret.Inst.ZoneThresholdNum = d.Get("zone_threshold_num").(int)
	ret.Inst.SrcPortRangeEnd = d.Get("src_port_range_end").(string)
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.LevelNum = d.Get("level_num").(string)
	ret.Inst.SrcPortRangeStart = d.Get("src_port_range_start").(string)
	return ret
}
