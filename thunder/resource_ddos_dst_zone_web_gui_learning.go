package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDstZoneWebGuiLearning() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_dst_zone_web_gui_learning`: Configure TPS Wizard Learning\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDstZoneWebGuiLearningCreate,
		UpdateContext: resourceDdosDstZoneWebGuiLearningUpdate,
		ReadContext:   resourceDdosDstZoneWebGuiLearningRead,
		DeleteContext: resourceDdosDstZoneWebGuiLearningDelete,

		Schema: map[string]*schema.Schema{
			"duration": {
				Type: schema.TypeString, Optional: true, Default: "6hour", Description: "'1minute': 1 minute; '6hour': 6 hours; '12hour': 12 hours; '24hour': 24 hours; '7day': 7 days;",
			},
			"starting_time": {
				Type: schema.TypeString, Optional: true, Description: "Configure learning starting time",
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
func resourceDdosDstZoneWebGuiLearningCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiLearningCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiLearning(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiLearningRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDstZoneWebGuiLearningUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiLearningUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiLearning(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDstZoneWebGuiLearningRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDstZoneWebGuiLearningDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiLearningDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiLearning(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDstZoneWebGuiLearningRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDstZoneWebGuiLearningRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDstZoneWebGuiLearning(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDstZoneWebGuiLearning(d *schema.ResourceData) edpt.DdosDstZoneWebGuiLearning {
	var ret edpt.DdosDstZoneWebGuiLearning
	ret.Inst.Duration = d.Get("duration").(string)
	ret.Inst.StartingTime = d.Get("starting_time").(string)
	//omit uuid
	ret.Inst.ZoneName = d.Get("zone_name").(string)
	return ret
}
