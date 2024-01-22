package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHarmonyControllerProfileReSync() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_harmony_controller_profile_re_sync`: re sync some options to harmony controller\n\n__PLACEHOLDER__",
		CreateContext: resourceHarmonyControllerProfileReSyncCreate,
		UpdateContext: resourceHarmonyControllerProfileReSyncUpdate,
		ReadContext:   resourceHarmonyControllerProfileReSyncRead,
		DeleteContext: resourceHarmonyControllerProfileReSyncDelete,

		Schema: map[string]*schema.Schema{
			"analytics_bus": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "re-sync analtyics bus connections",
			},
			"schema_registry": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "re-sync the schema registry",
			},
		},
	}
}
func resourceHarmonyControllerProfileReSyncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileReSyncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileReSync(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerProfileReSyncRead(ctx, d, meta)
	}
	return diags
}

func resourceHarmonyControllerProfileReSyncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileReSyncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileReSync(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHarmonyControllerProfileReSyncRead(ctx, d, meta)
	}
	return diags
}
func resourceHarmonyControllerProfileReSyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileReSyncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileReSync(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHarmonyControllerProfileReSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHarmonyControllerProfileReSyncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHarmonyControllerProfileReSync(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHarmonyControllerProfileReSync(d *schema.ResourceData) edpt.HarmonyControllerProfileReSync {
	var ret edpt.HarmonyControllerProfileReSync
	ret.Inst.AnalyticsBus = d.Get("analytics_bus").(int)
	ret.Inst.SchemaRegistry = d.Get("schema_registry").(int)
	return ret
}
