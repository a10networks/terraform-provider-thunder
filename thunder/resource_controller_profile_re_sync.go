package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceControllerProfileReSync() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_controller_profile_re_sync`: re sync some options to controller\n\n__PLACEHOLDER__",
		CreateContext: resourceControllerProfileReSyncCreate,
		UpdateContext: resourceControllerProfileReSyncUpdate,
		ReadContext:   resourceControllerProfileReSyncRead,
		DeleteContext: resourceControllerProfileReSyncDelete,

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
func resourceControllerProfileReSyncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileReSyncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileReSync(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerProfileReSyncRead(ctx, d, meta)
	}
	return diags
}

func resourceControllerProfileReSyncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileReSyncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileReSync(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceControllerProfileReSyncRead(ctx, d, meta)
	}
	return diags
}
func resourceControllerProfileReSyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileReSyncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileReSync(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceControllerProfileReSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceControllerProfileReSyncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointControllerProfileReSync(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointControllerProfileReSync(d *schema.ResourceData) edpt.ControllerProfileReSync {
	var ret edpt.ControllerProfileReSync
	ret.Inst.AnalyticsBus = d.Get("analytics_bus").(int)
	ret.Inst.SchemaRegistry = d.Get("schema_registry").(int)
	return ret
}
