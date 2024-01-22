package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceZoneLocalZoneCfg() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_zone_local_zone_cfg`: Local zone configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceZoneLocalZoneCfgCreate,
		UpdateContext: resourceZoneLocalZoneCfgUpdate,
		ReadContext:   resourceZoneLocalZoneCfgRead,
		DeleteContext: resourceZoneLocalZoneCfgDelete,

		Schema: map[string]*schema.Schema{
			"local_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set to local zone",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceZoneLocalZoneCfgCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneLocalZoneCfgCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZoneLocalZoneCfg(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceZoneLocalZoneCfgRead(ctx, d, meta)
	}
	return diags
}

func resourceZoneLocalZoneCfgUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneLocalZoneCfgUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZoneLocalZoneCfg(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceZoneLocalZoneCfgRead(ctx, d, meta)
	}
	return diags
}
func resourceZoneLocalZoneCfgDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneLocalZoneCfgDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZoneLocalZoneCfg(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceZoneLocalZoneCfgRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceZoneLocalZoneCfgRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointZoneLocalZoneCfg(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointZoneLocalZoneCfg(d *schema.ResourceData) edpt.ZoneLocalZoneCfg {
	var ret edpt.ZoneLocalZoneCfg
	ret.Inst.LocalType = d.Get("local_type").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
