package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneSrcPortZoneSrcPortLevel() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_src_port_zone_src_port_level`: Policy Level Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneSrcPortZoneSrcPortLevelCreate,
		UpdateContext: resourceDdosDstZoneSrcPortZoneSrcPortLevelUpdate,
		ReadContext:   resourceDdosDstZoneSrcPortZoneSrcPortLevelRead,
		DeleteContext: resourceDdosDstZoneSrcPortZoneSrcPortLevelDelete,

		Schema: map[string]*schema.Schema{
			"indicator_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type: schema.TypeString, Required: true, Description: "'pkt-rate': rate of incoming packets; 'bit-rate': rate of incoming bits;",
						},
						"zone_threshold_num": {
							Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the src-port",
						},
						"zone_threshold_large_num": {
							Type: schema.TypeInt, Optional: true, Description: "Threshold of the entire zone for the src-port",
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
				Type: schema.TypeString, Required: true, Description: "'0': Default policy level; '1': Policy level 1;",
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
			"port_num": {
				Type: schema.TypeString, Required: true, Description: "PortNum",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "Protocol",
			},
		},
	}
}
func resourceDdosDstZoneSrcPortZoneSrcPortLevelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortLevelCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortLevel(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcPortZoneSrcPortLevelRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneSrcPortZoneSrcPortLevelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortLevelUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortLevel(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneSrcPortZoneSrcPortLevelRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneSrcPortZoneSrcPortLevelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortLevelDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortLevel(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneSrcPortZoneSrcPortLevelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneSrcPortZoneSrcPortLevelRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneSrcPortZoneSrcPortLevel(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosDstZoneSrcPortZoneSrcPortLevelIndicatorList(d []interface{}) []edpt.DdosDstZoneSrcPortZoneSrcPortLevelIndicatorList {

	count1 := len(d)
	ret := make([]edpt.DdosDstZoneSrcPortZoneSrcPortLevelIndicatorList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosDstZoneSrcPortZoneSrcPortLevelIndicatorList
		oi.Type = in["type"].(string)
		oi.ZoneThresholdNum = in["zone_threshold_num"].(int)
		oi.ZoneThresholdLargeNum = in["zone_threshold_large_num"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDdosDstZoneSrcPortZoneSrcPortLevel(d *schema.ResourceData) edpt.DdosDstZoneSrcPortZoneSrcPortLevel {
	var ret edpt.DdosDstZoneSrcPortZoneSrcPortLevel
	ret.Inst.IndicatorList = getSliceDdosDstZoneSrcPortZoneSrcPortLevelIndicatorList(d.Get("indicator_list").([]interface{}))
	ret.Inst.LevelNum = d.Get("level_num").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	ret.Inst.PortNum = d.Get("port_num").(string)
	ret.Inst.Protocol = d.Get("protocol").(string)
	return ret
}
